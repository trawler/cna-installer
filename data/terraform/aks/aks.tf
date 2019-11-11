terraform {
  required_version = ">= 0.12.6"
  backend "azurerm" {
  }
}

provider "azurerm" {
  version = "1.35.0"
}

provider "local" {
  version = "1.3.0"
}

resource "azurerm_resource_group" "k8s" {
  name     = var.k8s_resource_group_name
  location = var.az_location
}

// Create a Standard Cluster (Don't use Cluster Autoscaler)
resource "azurerm_kubernetes_cluster" "k8s" {
  count = var.cluster_autoscaling ? 0 : 1

  name                = format("%s-%s", var.cluster_owner, var.k8s_cluster_name)
  location            = azurerm_resource_group.k8s.location
  resource_group_name = azurerm_resource_group.k8s.name
  kubernetes_version  = var.k8s_version
  dns_prefix          = var.cluster_owner

  linux_profile {
    admin_username = "ubuntu"
    ssh_key {
      key_data = file(var.public_key_file)
    }
  }

  agent_pool_profile {
    count           = local.agent_count
    name            = var.agent_pool_name
    os_disk_size_gb = var.agent_os_disk_size_gb
    os_type         = var.agent_os_type
    type            = local.agent_pool_type
    vm_size         = var.agent_vm_size
  }

  service_principal {
    client_id     = var.azure_client_id
    client_secret = var.azure_client_secret
  }

  tags = {
    Environment = "Production"
  }

  lifecycle {
    ignore_changes = ["agent_pool_profile[0].count"]
  }
}

// if enableAutoscaler is true, Create a Cluster using the Autoscaler
resource "azurerm_kubernetes_cluster" "k8s-autoscaler" {
  count = var.cluster_autoscaling ? 1 : 0

  name                = format("%s-%s", var.cluster_owner, var.k8s_cluster_name)
  location            = azurerm_resource_group.k8s.location
  resource_group_name = azurerm_resource_group.k8s.name
  kubernetes_version  = var.k8s_version
  dns_prefix          = var.cluster_owner

  linux_profile {
    admin_username = "ubuntu"
    ssh_key {
      key_data = file(var.public_key_file)
    }
  }

  agent_pool_profile {
    count = local.agent_count

    enable_auto_scaling = var.cluster_autoscaling
    max_count           = local.autoscaling_max_count
    min_count           = local.autoscaling_min_count

    name            = var.agent_pool_name
    os_disk_size_gb = var.agent_os_disk_size_gb
    os_type         = var.agent_os_type
    type            = local.agent_pool_type
    vm_size         = var.agent_vm_size
  }

  service_principal {
    client_id     = var.azure_client_id
    client_secret = var.azure_client_secret
  }

  tags = {
    Environment = "Production"
  }

  lifecycle {
    ignore_changes = ["agent_pool_profile[0].count"]
  }
}

resource "local_file" "kubectl" {
  content  = var.cluster_autoscaling ? azurerm_kubernetes_cluster.k8s-autoscaler[0].kube_config_raw : azurerm_kubernetes_cluster.k8s[0].kube_config_raw
  filename = "${path.cwd}/../logs/generated/auth/kubeconfig"
}

resource "local_file" "client_certificate" {
  content  = var.cluster_autoscaling ? azurerm_kubernetes_cluster.k8s-autoscaler[0].kube_config.0.client_certificate : azurerm_kubernetes_cluster.k8s[0].kube_config.0.client_certificate
  filename = "${path.cwd}/../logs/generated/auth/client.pem"
}

// Create Azure DNS
data "azurerm_dns_zone" "base_domain" {
  name = var.base_domain
}

resource "azurerm_dns_zone" "k8s" {
  name                = local.cluster_fqdn
  resource_group_name = azurerm_resource_group.k8s.name
}

resource "azurerm_dns_ns_record" "base_domain" {
  name                = replace(local.cluster_fqdn, ".${var.base_domain}", "")
  zone_name           = var.base_domain
  resource_group_name = data.azurerm_dns_zone.base_domain.resource_group_name
  ttl                 = 300

  records = azurerm_dns_zone.k8s.name_servers
}
