# AWS Configuration
variable "aws_region" {
  description = "AWS region for resources"
  type        = string
  default     = "us-east-1"  # JP Morgan typically uses us-east-1
}

variable "aws_profile" {
  description = "AWS profile to use (from PCL toolkit)"
  type        = string
  default     = "testkitclustercred"  # JP Morgan PCL profile name
}

# Cluster Configuration
variable "cluster_name" {
  description = "Name of the EKS cluster"
  type        = string
  default     = "swacd-demo-cluster"
}

variable "environment" {
  description = "Environment name (dev, staging, prod)"
  type        = string
  default     = "dev"
}

variable "owner" {
  description = "Owner/Team responsible for the cluster"
  type        = string
  default     = "DevOps-Team"
}

variable "kubernetes_version" {
  description = "Kubernetes version for the EKS cluster"
  type        = string
  default     = "1.28"
}

# Network Configuration - Using Default VPC
# No additional VPC variables needed as we're using the default VPC

variable "enable_nat_gateway" {
  description = "Enable NAT gateway for private subnets"
  type        = bool
  default     = true
}

# Removed single_nat_gateway variable - using default VPC

# Cluster Access
variable "cluster_endpoint_public_access" {
  description = "Enable public access to cluster endpoint"
  type        = bool
  default     = true
}

variable "cluster_endpoint_private_access" {
  description = "Enable private access to cluster endpoint"
  type        = bool
  default     = false
}

# Node Group Configuration
variable "node_instance_types" {
  description = "EC2 instance types for EKS managed node groups"
  type        = list(string)
  default     = ["t3.small"]
}

variable "node_group_min_size" {
  description = "Minimum number of nodes in the node group"
  type        = number
  default     = 1
}

variable "node_group_max_size" {
  description = "Maximum number of nodes in the node group"
  type        = number
  default     = 5
}

variable "node_group_desired_size" {
  description = "Desired number of nodes in the node group"
  type        = number
  default     = 2
}

variable "node_taints" {
  description = "Taints to apply to nodes"
  type = list(object({
    key    = string
    value  = string
    effect = string
  }))
  default = []
}

# SWACD Specific Configuration
variable "enable_swacd_security_rules" {
  description = "Enable additional security group rules for SWACD operator"
  type        = bool
  default     = true
}

# Add-ons Configuration
variable "enable_cluster_autoscaler" {
  description = "Enable cluster autoscaler for the EKS cluster"
  type        = bool
  default     = true
}

variable "enable_aws_load_balancer_controller" {
  description = "Enable AWS Load Balancer Controller"
  type        = bool
  default     = true
}

variable "enable_external_dns" {
  description = "Enable External DNS for automatic DNS management"
  type        = bool
  default     = false
}

variable "enable_cert_manager" {
  description = "Enable cert-manager for automatic SSL certificate management"
  type        = bool
  default     = true
}