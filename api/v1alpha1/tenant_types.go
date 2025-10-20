package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type EdgeProviderRef struct {
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}

// TenantSpec defines desired tenant configuration
type TenantSpec struct {
	DisplayName          string            `json:"displayName,omitempty"`
	Lob                  string            `json:"lob,omitempty"`
	Environment          string            `json:"environment,omitempty"`
	ContactEmail         string            `json:"contactEmail,omitempty"`
	EdgeServiceProviders []EdgeProviderRef `json:"edgeServiceProviders,omitempty"`
}

// TenantStatus defines observed state
type TenantStatus struct {
	ObservedGeneration int64              `json:"observedGeneration,omitempty"`
	Conditions         []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type Tenant struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TenantSpec   `json:"spec,omitempty"`
	Status            TenantStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
type TenantList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Tenant `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Tenant{}, &TenantList{})
}
