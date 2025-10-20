package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type OriginServiceSpec struct {
	Endpoint    string          `json:"endpoint"`
	HealthCheck HealthCheckSpec `json:"healthCheck,omitempty"`
	Auth        AuthSpec        `json:"auth,omitempty"`
}

type HealthCheckSpec struct {
	Path string `json:"path"`
}

type AuthSpec struct {
	TLS TLSConfig `json:"tls,omitempty"`
}

type TLSConfig struct {
	Enabled bool `json:"enabled"`
}

type OriginServiceStatus struct {
	Phase string `json:"phase,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type OriginService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OriginServiceSpec   `json:"spec,omitempty"`
	Status OriginServiceStatus `json:"status,omitempty"`
}

type OriginServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OriginService `json:"items"`
}
