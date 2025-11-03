# 🧭 SWACD Operator – Multi-Project Scaffold

This repository hosts the **Secure Web API and Content Delivery (SWACD)** control-plane operators for JPMC’s next-generation multi-tenant edge platform.  
It combines **KCP (Kubernetes Control Plane)** orchestration with multiple **Kubebuilder-based operators** for Cloudflare, Akamai, Hostmaster (DNS), and CKMS (certs/keys).

---

## 🚀 Current Status (Phase 1 Completed)

✅ **KCP deployed locally via Helm**  
✅ **Ingress configured with HTTPS (https://kcp.local)**  
✅ **Workspaces created (root → swacd → providers/lob/tenants)**  
✅ **Helm validation script (`check-swacd.sh`) working**  
✅ **Multi-project Kubebuilder scaffold created & pushed (`feature/multi-project-scaffold`)**

---

## 🧩 Repository Structure

```text
swacd-operator/
├── ucp/          # Unified Control Plane (KCP workspace orchestration)
├── swacd/        # Core SWACD CRDs (Tenant, EdgeRoute, OriginService, Provider)
├── cloudflare/   # Cloudflare provider operator
├── akamai/       # Akamai provider operator
├── hostmaster/   # DNS / zone management
├── ckms/         # Certificate & key management
└── charts/kcp/   # Helm chart for local KCP deployment
```

Each sub-project includes:
- `cmd/main.go`
- `config/` (RBAC, manager, kustomize)
- `Dockerfile`, `Makefile`, `PROJECT`, and `go.mod`
- `test/e2e/` scaffolds

---

## 🏗️ Local KCP Deployment (Helm)

```bash
# 1️⃣ Create working directory
mkdir ~/kcp-helm-demo && cd ~/kcp-helm-demo

# 2️⃣ Create Helm chart
helm create kcp
# Update Chart.yaml, values.yaml, deployment.yaml, service.yaml, ingress.yaml as per repo

# 3️⃣ Create KIND cluster
kind create cluster --name kcp-demo

# 4️⃣ Deploy KCP using Helm
helm upgrade --install my-kcp ./kcp -f ./kcp/values.yaml

# 5️⃣ Verify resources
kubectl get pods -A
kubectl get svc -A
kubectl get ingress -A

# 6️⃣ Add local DNS entry
sudo vim /etc/hosts
127.0.0.1 kcp.local
127.0.0.1 my-kcp.local

# 7️⃣ Access in browser
https://kcp.local   # should return 403 (expected)
```

---

## 🧠 Workspace Tree Validation

After port-forwarding:
```bash
kubectl port-forward svc/my-kcp 6443:6443 &
kubectl exec -it deployment/my-kcp -- ls /data
kubectl cp deployment/my-kcp:/data/admin.kubeconfig ./admin.kubeconfig
export KUBECONFIG=$(pwd)/admin.kubeconfig
kubectl ws tree
```

Expected tree:
```bash
.
└── root
    └── swacd
        ├── lob
        │   └── tenant-demo
        └── providers
            ├── akamai
            └── cloudflare
```

---

## 🧾 Environment Validation Script

`check-swacd.sh`
```bash
#!/bin/bash
echo "------------------------------------------"
echo " SWACD / KCP Environment Validation Script"
echo "------------------------------------------"

kubectl config current-context
kubectl get nodes -o wide
kubectl get pods -A -o wide
kubectl get svc -A
kubectl get ingress -A
helm list -A

echo
echo "Testing HTTPS endpoint (https://kcp.local/readyz):"
curl -sk https://kcp.local/readyz || echo "kcp.local not reachable"

if [ -f "./admin.kubeconfig" ]; then
  export KUBECONFIG=$(pwd)/admin.kubeconfig
  kubectl ws tree || echo "Unable to display workspace tree"
else
  echo "admin.kubeconfig not found in current directory"
fi

kubectl cluster-info || echo "Cluster info not available"
echo
echo "✅ KCP + Helm + Ingress + Workspaces check complete"
```

Run with:
```bash
chmod +x check-swacd.sh
./check-swacd.sh
```

---

## 🧱 Multi-Project Scaffold Creation

```bash
# 1️⃣ Create a feature branch
git checkout -b feature/multi-project-scaffold

# 2️⃣ Create sub-projects
mkdir -p ucp swacd cloudflare akamai hostmaster ckms

# 3️⃣ Initialize Kubebuilder for each
cd ucp
kubebuilder init --domain=internal.jpmc --owner 'JPMC'   --repo=github.com/Chalama7/swacd-operator/ucp

cd ../swacd
kubebuilder init --domain=swacd.internal.jpmc --owner 'JPMC'   --repo=github.com/Chalama7/swacd-operator/swacd

cd ../cloudflare
kubebuilder init --domain=cloudflare.internal.jpmc --owner 'JPMC'   --repo=github.com/Chalama7/swacd-operator/cloudflare

cd ../akamai
kubebuilder init --domain=akamai.internal.jpmc --owner 'JPMC'   --repo=github.com/Chalama7/swacd-operator/akamai

cd ../hostmaster
kubebuilder init --domain=hostmaster.internal.jpmc --owner 'JPMC'   --repo=github.com/Chalama7/swacd-operator/hostmaster

cd ../ckms
kubebuilder init --domain=ckms.internal.jpmc --owner 'JPMC'   --repo=github.com/Chalama7/swacd-operator/ckms
```

---

## 🧩 Git Workflow

```bash
# Commit & push
git add .
git commit -m "Scaffold: Multi-project Kubebuilder setup"
git push -u origin feature/multi-project-scaffold

# Open PR
# Title: Scaffold: Multi-project SWACD Operator
# Description: Includes 6 Kubebuilder sub-projects for modular controller design
```

---

## 🔮 Next Steps

1. **Define CRDs** inside `swacd/`:
   ```bash
   cd swacd
   kubebuilder create api --group swacd --version v1alpha1 --kind Tenant
   kubebuilder create api --group swacd --version v1alpha1 --kind EdgeRoute
   kubebuilder create api --group swacd --version v1alpha1 --kind OriginService
   kubebuilder create api --group swacd --version v1alpha1 --kind Provider
   ```

2. **Implement controllers**
   - Add reconciliation logic
   - Watch for changes to CRs
   - Trigger Cloudflare/Akamai APIs

3. **Deploy via Helm** to EKS once ready.

---

## 🧠 Authors
- **Chalama Reddy Venna** – SWACD Control Plane Engineer  
- **Cody (Architect)** – Technical guidance  
- **Team:** JPMC / Deloitte Cloud & Network Engineering  

---

_This document tracks the setup and scaffolding work up to KCP deployment and multi-operator scaffolding under SWACD project structure._
