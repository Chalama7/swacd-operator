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

package controller

import (
	"context"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"

	swacdv1alpha1 "github.com/Chalama7/swacd-operator/api/v1alpha1"
)

// CloudflareProviderReconciler reconciles a CloudflareProvider object
type CloudflareProviderReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=swacd.swacd.io,resources=cloudflareproviders,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=swacd.swacd.io,resources=cloudflareproviders/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=swacd.swacd.io,resources=cloudflareproviders/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the CloudflareProvider object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.22.1/pkg/reconcile
func (r *CloudflareProviderReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := logf.FromContext(ctx).WithValues("controller", "cloudflareprovider", "name", req.Name, "namespace", req.Namespace)

	var provider swacdv1alpha1.CloudflareProvider
	if err := r.Get(ctx, req.NamespacedName, &provider); err != nil {
		log.Error(err, "Failed to get CloudflareProvider")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	log.Info("Reconciling CloudflareProvider",
		"apiTokenSecretRef", provider.Spec.APITokenSecretRef,
		"zoneName", provider.Spec.ZoneName,
		"accountID", provider.Spec.AccountID,
	)

	provider.Status.Phase = "Active"
	provider.Status.Connected = true
	provider.Status.Message = "Simulated Cloudflare API connectivity successful"
	provider.Status.LastChecked = metav1.Now()
	provider.Status.Conditions = []metav1.Condition{
		{
			Type:               "Ready",
			Status:             metav1.ConditionTrue,
			Reason:             "Simulated",
			Message:            "Cloudflare API connectivity OK",
			LastTransitionTime: metav1.NewTime(time.Now()),
		},
	}

	if err := r.Status().Update(ctx, &provider); err != nil {
		log.Error(err, "Failed to update CloudflareProvider status")
		return ctrl.Result{}, err
	}

	log.Info("✅ Reconciled CloudflareProvider successfully")

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *CloudflareProviderReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&swacdv1alpha1.CloudflareProvider{}).
		Named("cloudflareprovider").
		Complete(r)
}
