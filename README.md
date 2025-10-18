# 🧭 SWACD Operator — Local Development Setup and Flow

This operator implements Kubernetes-style controllers for the **SWACD Control Plane PoC**, managing `Tenant` and `OriginService` custom resources.  
It supports full local reconciliation using `make`, `kubebuilder`, and `controller-runtime`.

---

## ⚙️ 1. Prerequisites

Make sure you have the following installed locally:

```bash
# Install dependencies
brew install go@1.24
brew install kubectl
brew install kind
brew install make
```

Then verify versions:

```bash
go version
kubectl version --client
make --version
```

---

## 🧱 2. Project Structure

```
swacd-operator/
├── api/
│   └── v1alpha1/
│       ├── tenant_types.go
│       ├── originservice_types.go
│       ├── groupversion_info.go
│
├── internal/
│   └── controller/
│       ├── tenant_controller.go
│       └── originservice_controller.go
│
├── config/
│   ├── crd/bases/
│   │   ├── swacd.swacd.io_tenants.yaml
│   │   └── swacd.swacd.io_originservices.yaml
│   ├── samples/
│   │   ├── swacd_v1alpha1_tenant.yaml
│   │   └── swacd_v1alpha1_originservice.yaml
│
└── cmd/
    └── main.go
```

---

## 🧩 3. Build and Install CRDs

Generate CRDs and RBAC configs:
```bash
make generate
make manifests
```

Install CRDs into your cluster:
```bash
make install
```

Verify CRDs:
```bash
kubectl get crds | grep swacd
```

You should see:
```
tenants.swacd.swacd.io
originservices.swacd.swacd.io
```

---

## 🚀 4. Run Controllers Locally

Start both controllers:
```bash
make run
```

You’ll see logs like:
```
INFO setup starting manager
INFO Starting Controller {"controller": "tenant"}
INFO Starting Controller {"controller": "originservice"}
```

---

## 🧠 5. Apply Sample CRDs

### Tenant
```bash
kubectl apply -f config/samples/swacd_v1alpha1_tenant.yaml
kubectl get tenants -o yaml
```

### OriginService
```bash
kubectl apply -f config/samples/swacd_v1alpha1_originservice.yaml --validate=false
kubectl get originservices -o yaml
```

---

## ✅ 6. Verify Reconciliation

When running correctly, your logs should show:

```
INFO  Tenant Spec details
INFO  ✅ Reconciled Tenant
INFO  🔍 OriginService Spec details
INFO  ✅ Reconciled OriginService
```

And your CRDs should reflect updated status fields:

```yaml
status:
  state: Active
  lastChecked: "2025-10-18T16:08:59-05:00"
  conditions:
  - type: Ready
    status: "True"
    reason: Reconciled
    message: OriginService originservice-sample successfully reconciled
```

---

## 🔁 7. Automatic Reconciliation on Startup

The operator automatically triggers reconciliation for existing CRs on startup (`main.go`):

```go
// Trigger reconciliation for existing OriginService CRs
go func() {
    time.Sleep(5 * time.Second)
    client := mgr.GetClient()
    var osList swacdv1alpha1.OriginServiceList
    client.List(context.Background(), &osList)
    for _, osvc := range osList.Items {
        osvc.Annotations["reconcile-trigger"] = time.Now().Format(time.RFC3339)
        client.Update(context.Background(), &osvc)
    }
}()
```

---

## 📦 8. Git Workflow

Commit and push your changes:
```bash
git add .
git commit -m "Working Tenant + OriginService controllers fully reconciled"
git push origin main
```

Ignore local KCP data:
```bash
echo ".kcp/" >> .gitignore
git rm -r --cached .kcp
git add .gitignore
git commit -m "Ignore local KCP data"
git push origin main
```

---

## 🧭 9. Next Steps

- [ ] Add **EdgeRoute** CRD + Controller  
- [ ] Add **Provider (Cloudflare / Akamai)** CRDs  
- [ ] Extend reconciliation logic to API integration  
- [ ] Document EKS + multi-cluster integration  

---

## 🧾 Credits

Developed and maintained by **Chalama Reddy Venna (Chalama7)**  
SWACD Control Plane | Deloitte | JPMC | 2025  
