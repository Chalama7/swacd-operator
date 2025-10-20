package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type CloudflareProviderSpec struct {
	AccountID         string    `json:"accountID"`
	APITokenSecretRef SecretRef `json:"apiTokenSecretRef"`
}

type SecretRef struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}

type CloudflareProviderStatus struct {
	Phase string `json:"phase,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type CloudflareProvider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudflareProviderSpec   `json:"spec,omitempty"`
	Status CloudflareProviderStatus `json:"status,omitempty"`
}

type CloudflareProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudflareProvider `json:"items"`
}
