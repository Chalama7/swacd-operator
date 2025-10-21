package controller

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"

	swacdv1alpha1 "github.com/Chalama7/swacd-operator/api/v1alpha1"
)

type AkamaiProviderReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=swacd.swacd.io,resources=akamaiproviders,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=swacd.swacd.io,resources=akamaiproviders/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=swacd.swacd.io,resources=akamaiproviders/finalizers,verbs=update

func (r *AkamaiProviderReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := logf.FromContext(ctx).WithValues("controller", "akamaiprovider", "controllerGroup", "swacd.swacd.io", "controllerKind", "AkamaiProvider")
	log.Info("🚀 Starting reconciliation for AkamaiProvider", "namespace", req.Namespace, "name", req.Name)

	var provider swacdv1alpha1.AkamaiProvider
	if err := r.Get(ctx, req.NamespacedName, &provider); err != nil {
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
	if err := r.Status().Update(ctx, &provider); err != nil {
		log.Error(err, "❌ Failed to update AkamaiProvider status")
		return ctrl.Result{}, err
	}

	log.Info("✅ Reconciled AkamaiProvider successfully", "name", provider.Name, "namespace", provider.Namespace)
	return ctrl.Result{}, nil
}

func (r *AkamaiProviderReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&swacdv1alpha1.AkamaiProvider{}).
		Named("akamaiprovider").
		Complete(r)
}
