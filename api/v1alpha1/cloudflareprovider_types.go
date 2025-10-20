package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CloudflareProviderSpec defines the desired state of CloudflareProvider
type CloudflareProviderSpec struct {
	APITokenSecretRef SecretRef `json:"apiTokenSecretRef"`
	ZoneName          string    `json:"zoneName"`
	AccountID         string    `json:"accountID"`
}

// CloudflareProviderStatus defines the observed state of CloudflareProvider
type CloudflareProviderStatus struct {
	State string `json:"state,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type CloudflareProvider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudflareProviderSpec   `json:"spec,omitempty"`
	Status CloudflareProviderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
type CloudflareProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudflareProvider `json:"items"`
}
