locals {
  agent_count           = var.cluster_autoscaling ? "1" : var.agent_count
  agent_pool_type       = var.cluster_autoscaling ? "VirtualMachineScaleSets" : "AvailabilitySet"
  autoscaling_min_count = var.cluster_autoscaling ? local.agent_count : var.agent_count
  autoscaling_max_count = var.cluster_autoscaling ? var.agent_count : local.autoscaling_min_count
}