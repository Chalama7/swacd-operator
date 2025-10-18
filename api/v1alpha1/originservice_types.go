/*
Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// OriginServiceSpec defines the desired state of OriginService
type OriginServiceSpec struct {
	Hostname        string `json:"hostname,omitempty"`
	Protocol        string `json:"protocol,omitempty"`
	Port            int32  `json:"port,omitempty"`
	HealthCheckPath string `json:"healthCheckPath,omitempty"`
}

// OriginServiceStatus defines the observed state of OriginService
type OriginServiceStatus struct {
	State       string             `json:"state,omitempty"`
	LastChecked string             `json:"lastChecked,omitempty"`
	Conditions  []metav1.Condition `json:"conditions,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// OriginService is the Schema for the originservices API
type OriginService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OriginServiceSpec   `json:"spec,omitempty"`
	Status OriginServiceStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// OriginServiceList contains a list of OriginService
type OriginServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OriginService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&OriginService{}, &OriginServiceList{})
}
