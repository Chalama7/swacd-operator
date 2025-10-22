#  SWACD Operator — Multicluster Controller Runtime

This repository contains the **SWACD Operator**, which defines and reconciles the core Custom Resources (CRDs) used in the **SWACD control plane POC**. 

**🚀 NEW**: This operator has been **converted to use multicluster controller-runtime** for cluster-aware resource management across multiple Kubernetes clusters.

It supports and reconciles the following resources:
-  **Tenant**
-  **OriginService**
-  **EdgeRoute**
-  **CloudflareProvider**
-  **AkamaiProvider**

## 🎯 Quick Start for New Users

**New to this repo?** See [MULTICLUSTER-SETUP.md](./MULTICLUSTER-SETUP.md) for complete setup instructions.

---

## Prerequisites

Ensure the following dependencies are installed **before running the operator**:

| Tool | Recommended Version | Install Command / Notes |
|------|---------------------|--------------------------|
| **Go** | `1.22+` | [Install Go](https://go.dev/doc/install) |
| **kubectl** | `1.28+` | `brew install kubectl` or follow [kubernetes.io/docs/tasks/tools](https://kubernetes.io/docs/tasks/tools/) |
| **Docker** | Latest | Required for running KIND and controller images |
| **KIND (Kubernetes in Docker)** | `v0.22+` | `go install sigs.k8s.io/kind@v0.22.0` |
| **KCP (Kubernetes Control Plane)** | `v0.23.0-alpha.1` | Download zip → unzip → move to PATH |
| **make** | default | Preinstalled on macOS/Linux |
| **git** | latest | `brew install git` or OS default |

---

## ⚙️ Step-by-Step Setup Guide

### 1️⃣ Clone the Repository
```bash
git clone https://github.com/Chalama7/swacd-operator.git
cd swacd-operator
```

### 2️⃣ Verify Go Dependencies
```bash
go mod tidy
```

### 3️⃣ Install Controller Tools (Kubebuilder utilities)
```bash
make install-tools
```
This installs the controller-gen binary at `bin/controller-gen`.

### 4️⃣ Generate CRDs and DeepCopy Code
```bash
make manifests
make generate
```
This will:
- Generate YAMLs under `config/crd/bases/`
- Update deep-copy files under `api/v1alpha1/`

### 5️⃣ Apply CRDs to Cluster (Local KIND or KCP)
If using **KCP**:
```bash
# Start KCP
./bin/kcp start
# Export KUBECONFIG
export KUBECONFIG=$(pwd)/.kcp/admin.kubeconfig
```
Then apply:
```bash
make install
```

Or for a local KIND cluster:
```bash
kind create cluster --name swacd-demo
kubectl apply -k config/crd/
```

### 6️⃣ Apply Sample Custom Resources
```bash
kubectl apply -k config/samples/
```
This installs demo objects:
- `acme-tenant`
- `acme-origin`
- `acme-edge-route`
- `cloudflare-prod`
- `akamai-prod`

### 7️⃣ Run the Controller
Run locally (without building container image):
```bash
make run
```
You should now see logs like:
```
✅ Reconciled Tenant successfully
✅ Reconciled OriginService successfully
✅ Reconciled EdgeRoute successfully
✅ Reconciled CloudflareProvider successfully
✅ Reconciled AkamaiProvider successfully
```

---

## 📂 Directory Structure
```
swacd-operator/
├── api/v1alpha1/              # CRD Go type definitions
├── internal/controller/       # Reconciler logic for each CRD
├── config/crd/bases/          # Generated CRDs
├── config/samples/            # Example CR instances
├── Makefile                   # Build + run automation
└── bin/                       # Controller-gen binaries
```

---

## 🧪 Verify Reconciliation
Run:
```bash
kubectl get tenants,originservices,edgeroutes,cloudflareproviders,akamaiproviders -A
```
Example output:
```
NAMESPACE   NAME                                AGE
default     tenant.swacd.swacd.io/acme-tenant   10m
default     originservice.swacd.swacd.io/acme-origin   10m
default     edgeroute.swacd.swacd.io/acme-edge-route   10m
default     cloudflareprovider.swacd.swacd.io/cloudflare-prod   10m
default     akamaiprovider.swacd.swacd.io/akamai-prod   10m
```

---

## 🧹 Cleanup
To remove all sample resources:
```bash
kubectl delete -k config/samples/
```
To delete the cluster:
```bash
kind delete cluster --name swacd-demo
```

---

## 🧾 Notes
- Tested on **macOS M4** and **Ubuntu 22.04**
- Works on **KCP workspaces** (`root:swacd`) and **local KIND clusters**
- All five CRDs reconcile locally and update status successfully
- Default branch: `main` (merged from `feature/status-phase`)
