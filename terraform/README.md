# SWACD Operator - EKS Terraform Deployment

This directory contains Terraform configuration for deploying a minimal EKS cluster for the SWACD operator in JP Morgan's sandbox environment using the PCL toolkit.

## Architecture

```
terraform/
├── main.tf          # Main infrastructure resources
├── variables.tf     # Configuration variables  
├── outputs.tf       # Cluster information outputs
└── README.md        # This documentation
```

## Features

- **JP Morgan PCL Integration**: Configured for `testkitclustercred` profile
- **Simplified Architecture**: Direct AWS resources, no complex modules
- **Federated User Compatible**: Minimal IAM configuration for sandbox users
- **Quick Deployment**: Single terraform apply for complete cluster
- **Cost Optimized**: t3.small instances, public subnets only
- **EKS 1.28**: Latest stable Kubernetes version

## Quick Start

### Prerequisites

1. **Terraform**: Version >= 1.0 installed
2. **AWS CLI**: Configured with PCL toolkit credentials
3. **PCL Toolkit**: Configured with `testkitclustercred` profile
4. **kubectl**: For cluster interaction

### 1. Initialize Terraform
```bash
cd terraform
terraform init
```

### 2. Review and Apply Configuration
```bash
# Review the plan
terraform plan

# Apply the configuration
terraform apply
```

### 3. Configure kubectl
```bash
# The output will show this command:
aws eks --region us-east-1 update-kubeconfig --name swacd-demo-cluster --profile testkitclustercred
```

### 4. Verify Cluster
```bash
kubectl get nodes
kubectl cluster-info
```

## Configuration Details

### Cluster Specifications
- **Name**: `swacd-demo-cluster`
- **Version**: Kubernetes 1.28
- **Region**: us-east-1
- **Instance Type**: t3.small
- **Node Count**: 2 (desired), 1-5 (scaling range)

### Network Configuration
- **VPC CIDR**: 10.0.0.0/16
- **Subnets**: 2 public subnets (10.0.1.0/24, 10.0.2.0/24)
- **Availability Zones**: us-east-1a, us-east-1b

### PCL Toolkit Integration
This deployment is specifically configured for JP Morgan's sandbox environment:
- Uses `testkitclustercred` AWS profile by default
- Simplified IAM configuration for federated users
- Direct AWS resource usage (no complex modules)

## Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `cluster_name` | EKS cluster name | `swacd-demo-cluster` |
| `cluster_version` | Kubernetes version | `1.28` |
| `region` | AWS region | `us-east-1` |
| `instance_types` | EC2 instance types | `["t3.small"]` |
| `desired_capacity` | Desired node count | `2` |
| `max_capacity` | Maximum node count | `5` |
| `min_capacity` | Minimum node count | `1` |
| `aws_profile` | AWS CLI profile | `testkitclustercred` |

## Outputs

After successful deployment, Terraform provides:
- Cluster endpoint and ARN
- kubectl configuration command
- VPC and subnet details
- Node group information

## Cleanup

To destroy the infrastructure:
```bash
terraform destroy
```

## Troubleshooting

### Common Issues

1. **Cluster Name Conflict**: If you get a ResourceInUseException, change the `cluster_name` variable to a unique value.

2. **IAM Permissions**: Ensure your PCL toolkit credentials have sufficient permissions for EKS, VPC, and IAM operations.

3. **Region Mismatch**: Verify that your AWS profile is configured for the correct region (us-east-1).

### Verification Commands
```bash
# Check AWS credentials
aws sts get-caller-identity --profile testkitclustercred

# List EKS clusters
aws eks list-clusters --profile testkitclustercred

# Check Terraform state
terraform show
```

## Next Steps

After the cluster is running:
1. Deploy the SWACD operator using kubectl
2. Configure ingress and load balancers as needed
3. Set up monitoring and logging
4. Deploy sample applications to test the operator

## Support

This configuration is designed specifically for JP Morgan's PCL toolkit sandbox environment. For production deployments, consider:
- Private subnets with NAT gateways
- Additional security groups
- OIDC provider for service accounts
- Cluster autoscaler and monitoring add-ons
- **Terraform Configuration**: Check this README and Terraform docs
- **EKS Issues**: Refer to AWS EKS documentation
- **SWACD Operator**: Check the main project documentation