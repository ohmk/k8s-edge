package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// EdgeNode is a specification for a EdgeNode resource
type EdgeNode struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EdgeNodeSpec   `json:"spec"`
	Status EdgeNodeStatus `json:"status"`
}

type EdgeNodeSpec struct {
	Pods []corev1.Pod `json:"pods"`
}

type EdgeNodeStatus struct {
	Phase        string      `json:"phase,omitempty"` // TODO: enum..
	LastSyncedAt metav1.Time `json:"last_synced_at"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EdgeNodeList is a list of EdgeNode resources
type EdgeNodeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []EdgeNode `json:"items"`
}
