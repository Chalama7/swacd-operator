package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type EdgeProviderRef struct {
	// Name of the provider resource
	Name string `json:"name"`
	// APIGroup of the provider resource (e.g., "swacd.swacd.io")
	APIGroup string `json:"apiGroup"`
	// Kind of the provider resource (e.g., "CloudflareProvider", "AkamaiProvider")
	Kind string `json:"kind"`
	// Namespace of the provider resource (optional, defaults to same namespace)
	Namespace string `json:"namespace,omitempty"`
}

// TenantSpec defines desired tenant configuration
type TenantSpec struct {
	DisplayName          string            `json:"displayName,omitempty"`
	Lob                  string            `json:"lob,omitempty"`
	Environment          string            `json:"environment,omitempty"`
	ContactEmail         string            `json:"contactEmail,omitempty"`
	EdgeServiceProviders []EdgeProviderRef `json:"edgeServiceProviders,omitempty"`
}

// TenantStatus defines the observed state of Tenant
type TenantStatus struct {
	Phase              string             `json:"phase,omitempty"` // Created, Updated, Failed
	ObservedGeneration int64              `json:"observedGeneration,omitempty"`
	Conditions         []metav1.Condition `json:"conditions,omitempty"`
	LastUpdated        metav1.Time        `json:"lastUpdated,omitempty"`
	Message            string             `json:"message,omitempty"`
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
