# Makefile for SWACD EKS Terraform Operations

TERRAFORM_DIR = .

# Colors for output
GREEN = \033[0;32m
YELLOW = \033[1;33m
RED = \033[0;31m
NC = \033[0m # No Color

.PHONY: help init plan apply destroy validate fmt lint clean status

help: ## Display this help message
	@echo "${GREEN}SWACD Operator EKS Terraform Commands${NC}"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  ${BLUE}%-18s${NC} %s\n", $$1, $$2}'
	@echo ""
	@echo "${GREEN}Examples:${NC}"
	@echo "  make pcl-help       # Show PCL toolkit usage"
	@echo "  make check-pcl      # Check PCL credentials status"
	@echo "  make check-tools    # Check if required tools are installed"
	@echo "  make init           # Initialize terraform"
	@echo "  make plan           # Plan terraform changes"  
	@echo "  make apply          # Apply terraform changes"
	@echo "  make kubeconfig     # Configure kubectl for the cluster"
	@echo "  make status         # Show cluster status"
	@echo "  make destroy        # Destroy the cluster"
	@echo ""

init: ## Initialize Terraform
	@echo "${GREEN}Initializing Terraform...${NC}"
	terraform init

plan: init ## Plan Terraform changes
	@echo "${YELLOW}Planning Terraform changes...${NC}"
	terraform plan -out=tfplan

apply: ## Apply Terraform changes
	@echo "${GREEN}Applying Terraform changes...${NC}"
	terraform apply tfplan

apply-auto: init ## Apply Terraform changes without confirmation
	@echo "${YELLOW}Auto-applying Terraform changes...${NC}"
	terraform apply -auto-approve

destroy: ## Destroy Terraform infrastructure
	@echo "${RED}WARNING: This will destroy the infrastructure!${NC}"
	@echo "Press Ctrl+C to cancel, or wait 10 seconds to continue..."
	@sleep 10
	terraform destroy

destroy-auto: ## Destroy Terraform infrastructure without confirmation
	@echo "${RED}Auto-destroying infrastructure...${NC}"
	terraform destroy -auto-approve

validate: ## Validate Terraform configuration
	@echo "${GREEN}Validating Terraform configuration...${NC}"
	terraform validate

fmt: ## Format Terraform files
	@echo "${GREEN}Formatting Terraform files...${NC}"
	terraform fmt -recursive .

lint: fmt validate ## Lint Terraform configuration
	@echo "${GREEN}Linting Terraform configuration...${NC}"
	@if command -v tflint >/dev/null 2>&1; then \
		tflint --recursive; \
	else \
		echo "${YELLOW}tflint not installed, skipping linting${NC}"; \
	fi

clean: ## Clean Terraform temporary files
	@echo "${GREEN}Cleaning Terraform temporary files...${NC}"
	find . -name "*.tfplan" -delete
	find . -name "*.tfstate.backup" -delete
	find . -name ".terraform.lock.hcl" -delete
	find . -type d -name ".terraform" -exec rm -rf {} + 2>/dev/null || true

status: ## Show Terraform state status
	@echo "${GREEN}Terraform state status:${NC}"
	terraform show

output: ## Show Terraform outputs
	@echo "${GREEN}Terraform outputs:${NC}"
	terraform output

refresh: ## Refresh Terraform state
	@echo "${GREEN}Refreshing Terraform state...${NC}"
	terraform refresh

# Kubernetes operations after cluster deployment
kubeconfig: ## Update kubectl configuration
	@echo "${GREEN}Updating kubectl configuration...${NC}"
	@CLUSTER_NAME=$$(terraform output -raw cluster_id 2>/dev/null); \
	AWS_REGION=$$(terraform output -raw cluster_info 2>/dev/null | jq -r '.region' 2>/dev/null || echo 'us-east-1'); \
	AWS_PROFILE=$$(terraform output -raw cluster_info 2>/dev/null | jq -r '.aws_profile' 2>/dev/null || echo 'adfs'); \
	if [ -n "$$CLUSTER_NAME" ]; then \
		aws eks --region $$AWS_REGION update-kubeconfig --name $$CLUSTER_NAME --profile $$AWS_PROFILE; \
		echo "${GREEN}kubectl configured for cluster: $$CLUSTER_NAME${NC}"; \
	else \
		echo "${RED}Could not get cluster name from terraform output${NC}"; \
	fi

deploy-swacd: kubeconfig ## Deploy SWACD operator to the cluster
	@echo "${GREEN}Deploying SWACD operator to cluster...${NC}"
	kubectl apply -f ../config/crd/bases/
	kubectl apply -f ../config/rbac/
	kubectl apply -f ../config/manager/
	@echo "${GREEN}SWACD operator deployed successfully${NC}"

# Quick deployment targets
quick-deploy: ## Quick deploy (plan + apply)
	$(MAKE) plan
	$(MAKE) apply

full-deploy: quick-deploy deploy-swacd ## Full deployment including SWACD operator

# Security and compliance
security-scan: ## Run security scanning on Terraform files
	@echo "${GREEN}Running security scan...${NC}"
	@if command -v tfsec >/dev/null 2>&1; then \
		tfsec .; \
	else \
		echo "${YELLOW}tfsec not installed, skipping security scan${NC}"; \
		echo "Install with: brew install tfsec or go install github.com/aquasecurity/tfsec/cmd/tfsec@latest"; \
	fi

cost-estimate: ## Estimate costs using infracost (if available)
	@echo "${GREEN}Estimating costs...${NC}"
	@if command -v infracost >/dev/null 2>&1; then \
		infracost breakdown --path .; \
	else \
		echo "${YELLOW}infracost not installed, skipping cost estimation${NC}"; \
		echo "Install from: https://www.infracost.io/docs/#quick-start"; \
	fi

# PCL Toolkit helpers
check-pcl: ## Check PCL credentials status
	@echo "${GREEN}Checking PCL credentials...${NC}"
	@aws sts get-caller-identity --profile testkitclustercred >/dev/null 2>&1 && \
		echo "${GREEN}PCL credentials are active${NC}" || \
		echo "${RED}PCL credentials expired or not found. Run: ./pcl aws --sandbox-user --domain <domain> --sid <sid> --profile-name testkitclustercred${NC}"

pcl-help: ## Show PCL toolkit usage
	@echo "${GREEN}PCL Toolkit Usage:${NC}"
	@echo ""
	@echo "1. Download PCL toolkit from AWS S3 console"
	@echo "2. Run: ./pcl aws --sandbox-user --domain <your-domain> --sid <your-sid> --profile-name testkitclustercred"
	@echo "3. Enter desktop password and RSA token"
	@echo "4. Verify: aws sts get-caller-identity --profile testkitclustercred"
	@echo ""
	@echo "Note: Credentials expire after 8 hours"

# Utility targets
check-tools: ## Check if required tools are installed
	@echo "${GREEN}Checking required tools...${NC}"
	@command -v terraform >/dev/null 2>&1 || (echo "${RED}terraform not found${NC}" && exit 1)
	@command -v aws >/dev/null 2>&1 || (echo "${RED}aws cli not found${NC}" && exit 1)
	@command -v kubectl >/dev/null 2>&1 || (echo "${RED}kubectl not found${NC}" && exit 1)
	@echo "${GREEN}All required tools are installed${NC}"
	@$(MAKE) check-pcl

version: ## Show tool versions
	@echo "${GREEN}Tool versions:${NC}"
	@echo "Terraform: $$(terraform version --json | jq -r '.terraform_version')"
	@echo "AWS CLI: $$(aws --version | cut -d' ' -f1)"
	@echo "kubectl: $$(kubectl version --client --short | cut -d' ' -f3)"