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
  default     = "francecentral"
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

variable "dns_prefix" {
  description = "(Required) DNS prefix specified when creating the managed cluster. Changing this forces a new resource to be created."
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

  default = 1
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

variable "argocd_workflow_repo_url" {
  description = "(Required) Source of the application workflow manifests"
  type        = "string"
}

variable "argocd_workflow_repo_ssh_key_path" {
  description = <<EOF
(Required) Path to the ssh private key file for the workflow repo,
relative to the platform/azure directory.
I.E, for files in the secrets folder,
the path should be ../../secrets/<filename>.
EOF

  type = "string"
}

variable "argocd_workflow_repo_path" {
  description = "(Required) Path to the application workflow manifests in the source repository."
  type        = "string"
  default     = ""
}

variable "argocd_root_app_repo_url" {
  description = "(Required) Path to the root application manifests / helm chart."
  type        = "string"
}

variable "argocd_root_app_repo_ssh_key_path" {
  description = <<EOF
(Required) Path to the ssh private key file for the root app repo,
relative to the platform/azure directory.
I.E, for files in the secrets folder,
the path should be ../../secrets/<filename>.
EOF

  type = "string"
}

variable "argocd_root_app_repo_path" {
  description = "(Required) Source of the root application manifest / helm chart."
  type        = "string"
  default     = ""
}

variable "harbor_admin_password" {
  description = "Admin password for the Harbor application"
  type        = "string"
  default     = ""
}

variable "harbor_db_host" {
  description = "External DB host for the Harbor application"
  type        = "string"
  default     = ""
}

variable "harbor_db_user" {
  description = "User for External DB host for the Harbor application"
  type        = "string"
  default     = ""
}

variable "harbor_db_password" {
  description = "Password for External DB host for the Harbor application"
  type        = "string"
  default     = ""
}

variable "patroni_users_password" {
  description = "Patroni password for 'postgres' 'admin' and 'standby' users"
  type        = "string"
  default     = ""
}

// dex oidc settings
variable "dex_oidc_client_id" {
  description = "(Required) Client ID for the Dex SSO Authentication"
  type        = "string"
}

variable "dex_oidc_client_secret" {
  description = "(Required) Client secret for the Dex SSO Authentication"
  type        = "string"
}

// istio Kiali
variable "istio_kiali_password" {
  type = "string"
}

variable "istio_kialli_username" {
  type = "string"
}

// istio Grafana
variable "istio_grafana_password" {
  type = "string"
}

variable "istio_grafana_username" {
  type = "string"
}
