variable "base_domain" {
  description = "(Required) DNS base domain to be used for the cluster's APIs"
  type        = "string"
}

variable "k8s_resource_group_name" {
  description = " (Required) Specifies the Resource Group where the Managed Kubernetes Cluster should exist. Changing this forces a new resource to be created."
  type        = "string"
}

variable "az_location" {
  description = "(Required) The location where the Managed Kubernetes Cluster should be created. Changing this forces a new resource to be created."
  type        = "string"
}

variable "k8s_cluster_name" {
  description = "(Required) The name of the Managed Kubernetes Cluster to create. Changing this forces a new resource to be created."
  type        = "string"
}

variable "k8s_version" {
  description = <<EOF
(Optional) Version of Kubernetes specified when creating the AKS managed cluster. If not specified,
the latest recommended version will be used at provisioning time (but won't auto-upgrade).
EOF
  type        = "string"
}

variable "cluster_owner" {
  description = "(Required) the user or group that created the cluster. Changing this forces a new resource to be created."
  type        = "string"
}

variable "public_key_file" {
  description = <<EOF
(required) the name of the SSH public key file to be provisioned from the secrets directory
as the SSH key for the 'ubuntu' user."
EOF
  type        = "string"
}

variable "agent_vm_size" {
  description = "(Required) The size of each VM in the Agent Pool (e.g. Standard_DS2_v2)."
  type        = "string"
}

variable "agent_count" {
  description = <<EOF
(Required) Number of Agents (VMs) in the Pool. Possible values must be in
the range of 1 to 100 (inclusive). Defaults to 1.
EOF
  type        = "string"
  default     = 1
}

variable "agent_os_type" {
  description = <<EOF
(Optional) The Operating System used for the Agents. Possible values are Linux and Windows.
Changing this forces a new resource to be created. Defaults to Linux.
EOF

  default = "Linux"
}

variable "agent_os_disk_size_gb" {
  description = "(Optional) The Agent Operating System disk size in GB. "
}

variable "agent_pool_name" {
  description = <<EOF
 (Required) Unique name of the Agent Pool Profile in the context of the Subscription and Resource Group. Changing this forces a new resource to be created.
EOF
}

variable "azure_client_id" {
  type        = "string"
  description = "(Required) The Client ID for the Service Principal."
}

variable "azure_client_secret" {
  type        = "string"
  description = "(Required) The Client Secret for the Service Principal."
}

variable "argocd_dex_liveness_readyness_path" {
  type    = "string"
  default = "/api/dex/.well-known/openid-configuration"
}
