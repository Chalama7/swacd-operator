/*
Copyright 2025.
*/

package controller

import (
	"context"
	"time"

	"k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"

	swacdv1alpha1 "github.com/Chalama7/swacd-operator/api/v1alpha1"

	mcbuilder "sigs.k8s.io/multicluster-runtime/pkg/builder"
	mcmanager "sigs.k8s.io/multicluster-runtime/pkg/manager"
	mcreconcile "sigs.k8s.io/multicluster-runtime/pkg/reconcile"
)

// EdgeRouteReconciler reconciles an EdgeRoute object
type EdgeRouteReconciler struct {
	mgr mcmanager.Manager
}

// +kubebuilder:rbac:groups=swacd.swacd.io,resources=edgeroutes,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=swacd.swacd.io,resources=edgeroutes/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=swacd.swacd.io,resources=edgeroutes/finalizers,verbs=update

func (r *EdgeRouteReconciler) Reconcile(ctx context.Context, req mcreconcile.Request) (ctrl.Result, error) {
	logger := logf.FromContext(ctx).WithValues("EdgeRoute", req.NamespacedName)
	logger.Info("🚀 Starting reconciliation for EdgeRoute")
	logger.Info("Detected cluster", "clusterName", req.ClusterName)

	cl, err := r.mgr.GetCluster(ctx, req.ClusterName)
	if err != nil {
		logger.Error(err, "Failed to get cluster")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	cc := cl.GetClient()

	// Fetch the EdgeRoute instance
	var route swacdv1alpha1.EdgeRoute
	if err := cc.Get(ctx, req.NamespacedName, &route); err != nil {
		logger.Error(err, "❌ Unable to fetch EdgeRoute")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// Print Spec details
	logger.Info("🔍 EdgeRoute Spec details",
		"fqdn", route.Spec.FQDN,
		"backendRefs", route.Spec.BackendRefs,
		"cacheEnabled", route.Spec.Cache.Enabled,
	)

	// Initialize status if empty
	if route.Status.Phase == "" {
		route.Status.Phase = "Pending"
		logger.Info("🕓 Setting initial status.phase to Pending")
	}

	// Simulate status update logic (e.g., after route validation)
	route.Status.Phase = "Active"
	route.Status.LastUpdated = v1.Now()

	// Update Conditions list
	condition := v1.Condition{
		Type:               "Configured",
		Status:             v1.ConditionTrue,
		Reason:             "ReconciledSuccessfully",
		Message:            "EdgeRoute successfully configured",
		LastTransitionTime: v1.Now(),
	}
	route.Status.Conditions = []v1.Condition{condition}

	// Apply status update
	if err := cc.Status().Update(ctx, &route); err != nil {
		logger.Error(err, "❌ Failed to update EdgeRoute status")
		return ctrl.Result{}, err
	}

	logger.Info("✅ Updated EdgeRoute status successfully",
		"phase", route.Status.Phase,
		"lastUpdated", route.Status.LastUpdated,
	)

	logger.Info("✅ Reconciled EdgeRoute successfully",
		"name", route.Name,
		"namespace", route.Namespace,
	)

	// Requeue every 30 seconds to simulate periodic sync
	return ctrl.Result{RequeueAfter: 30 * time.Second}, nil
}

func (r *EdgeRouteReconciler) SetupWithManager(mgr mcmanager.Manager) error {
	r.mgr = mgr
	return mcbuilder.
		ControllerManagedBy(mgr).
		For(&swacdv1alpha1.EdgeRoute{}).
		Named("edgeroute").
		Complete(r)
}
