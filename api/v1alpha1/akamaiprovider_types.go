package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type AkamaiProviderSpec struct {
	ContractID          string    `json:"contractID"`
	CredentialSecretRef SecretRef `json:"credentialSecretRef"`
}

type AkamaiProviderStatus struct {
	Phase string `json:"phase,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type AkamaiProvider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AkamaiProviderSpec   `json:"spec,omitempty"`
	Status AkamaiProviderStatus `json:"status,omitempty"`
}

type AkamaiProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AkamaiProvider `json:"items"`
}
