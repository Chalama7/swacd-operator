package v1alpha1

// SecretRef defines a reference to a Kubernetes Secret key.
type SecretRef struct {
	// Name of the Secret
	Name string `json:"name"`
	// Key inside the Secret data
	Key string `json:"key"`
}

// +kubebuilder:object:generate=true
// This is a common type, not a top-level CRD.
type ProviderCommonSpec struct {
	// Reference to Secret containing provider credentials
	CredentialRef SecretRef `json:"credentialRef"`
}
