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
	"fmt"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"

	swacdv1alpha1 "github.com/Chalama7/swacd-operator/api/v1alpha1"

	mcbuilder "sigs.k8s.io/multicluster-runtime/pkg/builder"
	mcmanager "sigs.k8s.io/multicluster-runtime/pkg/manager"
	mcreconcile "sigs.k8s.io/multicluster-runtime/pkg/reconcile"
)

// OriginServiceReconciler reconciles a OriginService object
type OriginServiceReconciler struct {
	mgr mcmanager.Manager
}

// +kubebuilder:rbac:groups=swacd.swacd.io,resources=originservices,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=swacd.swacd.io,resources=originservices/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=swacd.swacd.io,resources=originservices/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the OriginService object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.22.1/pkg/reconcile
func (r *OriginServiceReconciler) Reconcile(ctx context.Context, req mcreconcile.Request) (ctrl.Result, error) {
	log := logf.FromContext(ctx)
	log.Info("Detected cluster", "clusterName", req.ClusterName)

	log.Info("🚀 Starting reconciliation for OriginService", "name", req.NamespacedName)

	defer func() {
		log.Info("✅ Finished reconciliation for OriginService", "name", req.NamespacedName)
	}()

	cl, err := r.mgr.GetCluster(ctx, req.ClusterName)
	if err != nil {
		log.Error(err, "Failed to get cluster")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	cc := cl.GetClient()

	var originService swacdv1alpha1.OriginService
	if err := cc.Get(ctx, req.NamespacedName, &originService); err != nil {
		log.Error(err, "❌ Failed to get OriginService")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	log.Info("🔍 OriginService Spec details",
		"hostname", originService.Spec.Hostname,
		"protocol", originService.Spec.Protocol,
		"port", originService.Spec.Port,
		"healthCheckPath", originService.Spec.HealthCheckPath,
	)

	log.Info("ℹ️ OriginService Current Status",
		"state", originService.Status.State,
		"lastChecked", originService.Status.LastChecked,
		"conditionsCount", len(originService.Status.Conditions),
	)

	if originService.Spec.Hostname == "" || originService.Spec.HealthCheckPath == "" {
		originService.Status.State = "Pending"
	} else {
		originService.Status.State = "Active"
	}

	originService.Status.LastChecked = time.Now().Format(time.RFC3339)

	// Remove existing Ready conditions to avoid infinite appends
	filteredConditions := []metav1.Condition{}
	for _, cond := range originService.Status.Conditions {
		if cond.Type != "Ready" {
			filteredConditions = append(filteredConditions, cond)
		}
	}

	condition := metav1.Condition{
		Type:               "Ready",
		Status:             metav1.ConditionTrue,
		LastTransitionTime: metav1.Now(),
		Reason:             "Reconciled",
		Message:            fmt.Sprintf("OriginService %s successfully reconciled", originService.Name),
	}
	filteredConditions = append(filteredConditions, condition)
	originService.Status.Conditions = filteredConditions

	if err := cc.Status().Update(ctx, &originService); err != nil {
		log.Error(err, "❌ Failed to update OriginService status")
		return ctrl.Result{}, err
	}

	log.Info("✅ Reconciled OriginService successfully",
		"name", originService.Name,
		"state", originService.Status.State,
		"lastChecked", originService.Status.LastChecked,
	)

	return ctrl.Result{RequeueAfter: 30 * time.Second}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *OriginServiceReconciler) SetupWithManager(mgr mcmanager.Manager) error {
	r.mgr = mgr
	return mcbuilder.
		ControllerManagedBy(mgr).
		For(&swacdv1alpha1.OriginService{}).
		Named("originservice").
		Complete(r)
}
