package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TenantSpec defines the desired state of Tenant
type TenantSpec struct {
	// Line of Business or owning team
	Lob string `json:"lob,omitempty"`

	// Environment (e.g., dev, stage, prod)
	Environment string `json:"environment,omitempty"`

	// Contact email for notifications
	ContactEmail string `json:"contactEmail,omitempty"`
}

// TenantStatus defines the observed state of Tenant
type TenantStatus struct {
	// ObservedGeneration tracks the last reconciled spec generation
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// Conditions represents the latest available observations of an object's state
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Tenant is the Schema for the tenants API
type Tenant struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TenantSpec   `json:"spec,omitempty"`
	Status TenantStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TenantList contains a list of Tenant
type TenantList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Tenant `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Tenant{}, &TenantList{})
}
