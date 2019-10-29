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

variable "cluster_owner" {
  description = "(Required) DNS prefix specified when creating the managed cluster. Changing this forces a new resource to be created."
  type        = "string"
}
