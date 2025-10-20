package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EdgeRouteSpec defines the desired state of EdgeRoute
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

// EdgeRouteStatus defines the observed state of EdgeRoute
type EdgeRouteStatus struct {
	// Phase indicates the lifecycle phase (e.g., Pending, Active, Failed)
	Phase string `json:"phase,omitempty"`

	// Conditions represent the latest available observations of EdgeRoute state
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// LastUpdated shows when the route status was last updated
	LastUpdated metav1.Time `json:"lastUpdated,omitempty"`
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
