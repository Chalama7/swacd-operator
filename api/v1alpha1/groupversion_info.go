// +kubebuilder:object:generate=true
// +groupName=swacd.swacd.io
package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

var (
	GroupVersion = schema.GroupVersion{Group: "swacd.swacd.io", Version: "v1alpha1"}

	SchemeBuilder = &scheme.Builder{GroupVersion: GroupVersion}
	AddToScheme   = SchemeBuilder.AddToScheme
)

func init() {
	SchemeBuilder.Register(
		&Tenant{}, &TenantList{},
		&OriginService{}, &OriginServiceList{},
		&EdgeRoute{}, &EdgeRouteList{},
		&CloudflareProvider{}, &CloudflareProviderList{},
		&AkamaiProvider{}, &AkamaiProviderList{},
	)
}
