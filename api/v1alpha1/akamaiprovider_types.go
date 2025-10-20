package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AkamaiProviderSpec defines the desired state of AkamaiProvider
type AkamaiProviderSpec struct {
	CredentialSecretRef SecretRef `json:"credentialSecretRef"`
	ConfigSection       string    `json:"configSection"`
}

// AkamaiProviderStatus defines the observed state of AkamaiProvider
type AkamaiProviderStatus struct {
	State string `json:"state,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type AkamaiProvider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AkamaiProviderSpec   `json:"spec,omitempty"`
	Status AkamaiProviderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
type AkamaiProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AkamaiProvider `json:"items"`
}
