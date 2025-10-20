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

// tenant_controller.go contains the controller logic for Tenant resource reconciliation.

package controller

import (
	"context"
	"fmt"
	"time"

	swacdv1alpha1 "github.com/Chalama7/swacd-operator/api/v1alpha1"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

// TenantReconciler reconciles a Tenant object
type TenantReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// SetupWithManager sets up the controller with the Manager.
func (r *TenantReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&swacdv1alpha1.Tenant{}).
		Complete(r)
}

// SetupTenantReconciler registers the Tenant controller with the manager.
func SetupTenantReconciler(mgr manager.Manager) error {
	if err := swacdv1alpha1.AddToScheme(mgr.GetScheme()); err != nil {
		return err
	}
	reconciler := &TenantReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
	}
	return reconciler.SetupWithManager(mgr)
}

// +kubebuilder:rbac:groups=swacd.swacd.io,resources=tenants,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=swacd.swacd.io,resources=tenants/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=swacd.swacd.io,resources=tenants/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *TenantReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	tenant := &swacdv1alpha1.Tenant{}
	if err := r.Get(ctx, req.NamespacedName, tenant); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	tenant.Status.ObservedGeneration = tenant.Generation
	meta.SetStatusCondition(&tenant.Status.Conditions, metav1.Condition{
		Type:    "Ready",
		Status:  metav1.ConditionTrue,
		Reason:  "Reconciled",
		Message: fmt.Sprintf("Tenant %s successfully reconciled", tenant.Name),
	})

	// Log Tenant spec details for demo visibility
	log.Info("Tenant Spec details", "Lob", tenant.Spec.Lob, "Environment", tenant.Spec.Environment, "ContactEmail", tenant.Spec.ContactEmail)

	// ✅ Update Tenant Status
	tenant.Status.Phase = "Created"
	tenant.Status.Message = "Tenant successfully reconciled"
	tenant.Status.ObservedGeneration = tenant.Generation
	tenant.Status.LastUpdated = metav1.Now()

	if err := r.Status().Update(ctx, tenant); err != nil {
		log.Error(err, "❌ Failed to update Tenant status")
	} else {
		log.Info("✅ Updated Tenant status successfully")
	}

	if err := r.Status().Update(ctx, tenant); err != nil {
		log.Error(err, "unable to update Tenant status")
		return ctrl.Result{}, err
	}

	log.Info("✅ Reconciled Tenant", "name", tenant.Name)
	return ctrl.Result{RequeueAfter: 30 * time.Second}, nil
}
