terraform {
  required_version = ">= 0.12.6"
  backend "azurerm" {
  }
}

provider "azurerm" {
  version = "1.31.0"
}

provider "local" {
  version = "1.3.0"
}

resource "azurerm_resource_group" "k8s" {
  name     = var.k8s_resource_group_name
  location = var.az_location
}

resource "azurerm_kubernetes_cluster" "k8s" {
  name                = var.k8s_cluster_name
  location            = azurerm_resource_group.k8s.location
  resource_group_name = azurerm_resource_group.k8s.name
  kubernetes_version  = var.k8s_version
  dns_prefix          = var.dns_prefix

  linux_profile {
    admin_username = "ubuntu"
    ssh_key {
      key_data = file(var.public_key_file)
    }
  }

  agent_pool_profile {
    count           = var.agent_count
    name            = var.k8s_cluster_name
    vm_size         = var.agent_vm_size
    os_type         = var.agent_os_type
    os_disk_size_gb = var.agent_os_disk_size_gb
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
  content  = azurerm_kubernetes_cluster.k8s.kube_config_raw
  filename = "${path.module}/generated/auth/kubeconfig"
}

resource "local_file" "client_certificate" {
  content  = azurerm_kubernetes_cluster.k8s.kube_config.0.client_certificate
  filename = "${path.module}/generated/auth/client.pem"
}
