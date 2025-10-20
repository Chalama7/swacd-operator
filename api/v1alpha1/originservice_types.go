package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type OriginServiceSpec struct {
	Hostname        string `json:"hostname,omitempty"`
	Protocol        string `json:"protocol,omitempty"`
	Port            int    `json:"port,omitempty"`
	HealthCheckPath string `json:"healthCheckPath,omitempty"`
}

type OriginServiceStatus struct {
	State       string             `json:"state,omitempty"`
	LastChecked string             `json:"lastChecked,omitempty"`
	Conditions  []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type OriginService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OriginServiceSpec   `json:"spec,omitempty"`
	Status            OriginServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
type OriginServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OriginService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&OriginService{}, &OriginServiceList{})
}
