output "cluster_id" {
  description = "EKS cluster ID"
  value       = aws_eks_cluster.main.name
}

output "cluster_arn" {
  description = "EKS cluster ARN"
  value       = aws_eks_cluster.main.arn
}

output "cluster_endpoint" {
  description = "Endpoint for EKS control plane"
  value       = aws_eks_cluster.main.endpoint
}

output "cluster_security_group_id" {
  description = "Security group ID attached to the EKS cluster"
  value       = aws_eks_cluster.main.vpc_config[0].cluster_security_group_id
}

output "cluster_certificate_authority_data" {
  description = "Base64 encoded certificate data required to communicate with the cluster"
  value       = aws_eks_cluster.main.certificate_authority[0].data
}

output "cluster_version" {
  description = "The Kubernetes version for the EKS cluster"
  value       = aws_eks_cluster.main.version
}

output "vpc_id" {
  description = "ID of the VPC where the EKS cluster is deployed"
  value       = aws_vpc.main.id
}

output "subnets" {
  description = "List of IDs of subnets used by EKS"
  value       = aws_subnet.public[*].id
}

output "subnet_details" {
  description = "Details of subnets used by EKS"
  value       = "Using existing VPC subnets"
}

output "node_groups" {
  description = "EKS node groups"
  value = {
    main = {
      node_group_arn    = aws_eks_node_group.main.arn
      node_group_status = aws_eks_node_group.main.status
      instance_types    = aws_eks_node_group.main.instance_types
      scaling_config    = aws_eks_node_group.main.scaling_config
    }
  }
}

output "cluster_oidc_issuer_url" {
  description = "The URL on the EKS cluster for the OpenID Connect identity provider"
  value       = aws_eks_cluster.main.identity[0].oidc[0].issuer
}

output "oidc_provider_arn" {
  description = "The ARN of the OIDC Provider (requires OIDC provider to be created separately)"
  value       = "Not created in this simple configuration"
}

# Useful commands
output "kubectl_config_command" {
  description = "Command to configure kubectl"
  value       = "aws eks --region ${var.aws_region} update-kubeconfig --name ${aws_eks_cluster.main.name} --profile ${var.aws_profile}"
}

output "cluster_info" {
  description = "EKS cluster information summary"
  value = {
    name        = aws_eks_cluster.main.name
    endpoint    = aws_eks_cluster.main.endpoint
    version     = aws_eks_cluster.main.version
    region      = var.aws_region
    vpc_id      = aws_vpc.main.id
    environment = var.environment
  }
}