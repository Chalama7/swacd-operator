package controller

import (
	"context"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"

	swacdv1alpha1 "github.com/Chalama7/swacd-operator/api/v1alpha1"

	mcbuilder "sigs.k8s.io/multicluster-runtime/pkg/builder"
	mcmanager "sigs.k8s.io/multicluster-runtime/pkg/manager"
	mcreconcile "sigs.k8s.io/multicluster-runtime/pkg/reconcile"
)

type AkamaiProviderReconciler struct {
	mgr mcmanager.Manager
}

// +kubebuilder:rbac:groups=swacd.swacd.io,resources=akamaiproviders,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=swacd.swacd.io,resources=akamaiproviders/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=swacd.swacd.io,resources=akamaiproviders/finalizers,verbs=update

func (r *AkamaiProviderReconciler) Reconcile(ctx context.Context, req mcreconcile.Request) (ctrl.Result, error) {
	log := logf.FromContext(ctx).WithValues("controller", "akamaiprovider", "controllerGroup", "swacd.swacd.io", "controllerKind", "AkamaiProvider")
	log.Info("🚀 Starting reconciliation for AkamaiProvider", "namespace", req.Namespace, "name", req.Name, "cluster", req.ClusterName)

	// Get cluster-specific client
	cl, err := r.mgr.GetCluster(ctx, req.ClusterName)
	if err != nil {
		log.Error(err, "❌ Failed to get cluster")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	cc := cl.GetClient()

	var provider swacdv1alpha1.AkamaiProvider
	if err := cc.Get(ctx, req.NamespacedName, &provider); err != nil {
		log.Error(err, "❌ Failed to get AkamaiProvider")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	log.Info("🔍 AkamaiProvider Spec details",
		"baseURL", provider.Spec.BaseURL,
		"contractID", provider.Spec.ContractID,
		"groupID", provider.Spec.GroupID,
		"clientTokenSecretRef", provider.Spec.ClientTokenSecretRef,
		"clientSecretSecretRef", provider.Spec.ClientSecretSecretRef,
		"accessTokenSecretRef", provider.Spec.AccessTokenSecretRef,
	)

	provider.Status.State = "Active"
	if err := cc.Status().Update(ctx, &provider); err != nil {
		log.Error(err, "❌ Failed to update AkamaiProvider status")
		return ctrl.Result{}, err
	}

	log.Info("✅ Reconciled AkamaiProvider successfully", "name", provider.Name, "namespace", provider.Namespace, "cluster", req.ClusterName)
	return ctrl.Result{}, nil
}

func (r *AkamaiProviderReconciler) SetupWithManager(mgr mcmanager.Manager) error {
	r.mgr = mgr
	return mcbuilder.
		ControllerManagedBy(mgr).
		For(&swacdv1alpha1.AkamaiProvider{}).
		Complete(r)
}
