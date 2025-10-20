package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type EdgeRouteSpec struct {
	FQDN        string        `json:"fqdn"`
	BackendRefs []BackendRef  `json:"backendRefs,omitempty"`
	Cache       CacheSettings `json:"cache,omitempty"`
}

type BackendRef struct {
	Name string `json:"name"`
}

type CacheSettings struct {
	Enabled bool `json:"enabled"`
}

type EdgeRouteStatus struct {
	Phase string `json:"phase,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type EdgeRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EdgeRouteSpec   `json:"spec,omitempty"`
	Status            EdgeRouteStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
type EdgeRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EdgeRoute `json:"items"`
}
