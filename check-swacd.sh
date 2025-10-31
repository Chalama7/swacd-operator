#!/bin/bash

echo "------------------------------------------"
echo " SWACD / KCP Environment Validation Script"
echo "------------------------------------------"

echo
echo "Context:"
kubectl config current-context

echo
echo "Nodes:"
kubectl get nodes -o wide

echo
echo "Pods (all namespaces):"
kubectl get pods -A -o wide

echo
echo "Services:"
kubectl get svc -A

echo
echo "Ingress:"
kubectl get ingress -A

echo
echo "Helm Releases:"
helm list -A

echo
echo "Testing HTTPS endpoint (https://kcp.local/readyz):"
curl -sk https://kcp.local/readyz || echo "kcp.local not reachable"

echo
echo "Workspaces:"
if [ -f "./admin.kubeconfig" ]; then
  export KUBECONFIG=$(pwd)/admin.kubeconfig
  kubectl ws tree || echo "Unable to display workspace tree"
else
  echo "admin.kubeconfig not found in current directory"
fi

echo
echo "Cluster Info:"
kubectl cluster-info || echo "Cluster info not available"

echo

echo " KCP + Helm + Ingress + Workspaces check complete"
