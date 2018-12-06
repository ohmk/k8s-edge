package manager

import (
	"fmt"
	"time"

	"github.com/ohmk/k8s-edge/pkg/apis/edge/v1alpha1"
	clientset "github.com/ohmk/k8s-edge/pkg/client/clientset/versioned"
	informers "github.com/ohmk/k8s-edge/pkg/client/informers/externalversions"
	listers "github.com/ohmk/k8s-edge/pkg/client/listers/edge/v1alpha1"
	log "github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/workqueue"
)

type edgeController struct {
	kubeclientset   kubernetes.Interface
	edgeclientset   clientset.Interface
	queue           workqueue.RateLimitingInterface
	lister          listers.EdgeNodeLister
	synced          cache.InformerSynced
	informerFactory informers.SharedInformerFactory
}

func newEdgeController(masterURL, path string) (*edgeController, error) {
	config, err := clientcmd.BuildConfigFromFlags(masterURL, path)
	if err != nil {
		log.Fatalf("Failed to build config from: %s, %s", masterURL, path)
		return nil, err
	}
	kubeclientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Failed to build config from: %s, %s", masterURL, path)
		return nil, err
	}
	edgeclientset, err := clientset.NewForConfig(config)
	if err != nil {
		log.Fatalf("Failed to build config from: %s, %s", masterURL, path)
		return nil, err
	}

	informerFactory := informers.NewSharedInformerFactory(edgeclientset, time.Second*30)
	informer := informerFactory.Edge().V1alpha1().EdgeNodes()
	queue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())

	informer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(obj)
			if err != nil {
				runtime.HandleError(err)
				return
			}
			queue.Add(key)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(newObj)
			if err != nil {
				runtime.HandleError(err)
				return
			}
			queue.Add(key)
		},
		DeleteFunc: func(obj interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(obj)
			if err != nil {
				runtime.HandleError(err)
				return
			}
			queue.Add(key)
		},
	})

	return &edgeController{
		kubeclientset:   kubeclientset,
		edgeclientset:   edgeclientset,
		queue:           queue,
		lister:          informer.Lister(),
		synced:          informer.Informer().HasSynced,
		informerFactory: informerFactory,
	}, nil
}

func (e *edgeController) Run(stopCh <-chan struct{}, errCh chan<- error) {
	defer runtime.HandleCrash()
	defer e.queue.ShutDown()

	go e.informerFactory.Start(stopCh)

	log.Info("Waiting for the informer cache to sync")
	if !cache.WaitForCacheSync(stopCh, e.synced) {
		runtime.HandleError(fmt.Errorf("Timed out waiting for caches to sync"))
		return
	}

	log.Info("Starting workers")
	go wait.Until(e.runWorker, time.Second, stopCh)

	<-stopCh
	log.Info("Shutting down workers")
}

func (e *edgeController) runWorker() {
	for e.processNextWorkItem() {
	}
}

func (e *edgeController) processNextWorkItem() bool {
	obj, shutdown := e.queue.Get()
	if shutdown {
		return false
	}
	defer e.queue.Done(obj)

	key, ok := obj.(string)
	if !ok {
		log.Infof("Failed to get key from '%v'", obj)
		return true
	}

	err := e.syncHandler(key)
	if err == nil {
		e.queue.Forget(obj)
		return true
	}
	runtime.HandleError(fmt.Errorf("Sync %v failed with %v", key, err))
	e.queue.AddRateLimited(obj)

	return true
}

func (e *edgeController) syncHandler(key string) error {
	namespace, name, err := cache.SplitMetaNamespaceKey(key)

	edgeNode, err := e.lister.EdgeNodes(namespace).Get(name)
	if err != nil {
		log.Info(err)
		if errors.IsNotFound(err) {
			log.Infof("Edge node is deleted: %s", key)
			return nil
		}
	}

	for _, pod := range edgeNode.Spec.Pods {
		// TODO: Pod's namespace is always same as edgenode's?
		_, err := e.kubeclientset.CoreV1().Pods(namespace).Get(pod.ObjectMeta.Name+"-"+name, metav1.GetOptions{})
		if err != nil {
			if errors.IsNotFound(err) {
				log.Infof("%s/%s is not found", namespace, pod.ObjectMeta.Name+"-"+name)
				e.updateStatus(edgeNode, "Pending")
				return nil
			}
		}
	}

	// TODO FIXME:
	// Find a pod which is not existed in the EdgeNode.Spec
	pods, err := e.kubeclientset.CoreV1().Pods(namespace).List(metav1.ListOptions{})
	for _, pod := range pods.Items {
		if pod.Spec.NodeName != name {
			continue
		}
		ok := false
		for _, i := range edgeNode.Spec.Pods {
			if pod.ObjectMeta.Name == i.ObjectMeta.Name+"-"+name {
				ok = true
				break
			}
		}

		if !ok {
			e.updateStatus(edgeNode, "Pending")
			return nil
		}
	}
	e.updateStatus(edgeNode, "Running")
	return nil
}

func (e *edgeController) updateStatus(edgeNode *v1alpha1.EdgeNode, phase string) {

	if edgeNode.Status.Phase == phase {
		log.Infof("%s phase is already %s", edgeNode.ObjectMeta.Name, phase)
		return
	}

	edgeNode.Status.Phase = phase
	_, err := e.edgeclientset.EdgeV1alpha1().EdgeNodes(edgeNode.ObjectMeta.Namespace).Update(edgeNode)
	if err != nil {
		log.Warn(err)
	}
}
