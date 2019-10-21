provider "azurerm" {
  version = "1.35.0"
}

resource "azurerm_resource_group" "k8s" {
  name     = var.k8s_resource_group_name
  location = var.az_location
}


// Azure terraform backend storage account
resource "azurerm_storage_account" "tf-backend" {
  name                     = format("%stfstorage", var.dns_prefix)
  resource_group_name      = azurerm_resource_group.k8s.name
  location                 = azurerm_resource_group.k8s.location
  account_tier             = "Standard"
  account_replication_type = "LRS"

  tags = {
    environment = "dev"
  }
}

resource "azurerm_storage_container" "tf-storage-container" {
  name                  = "terraform-tfstate"
  storage_account_name  = azurerm_storage_account.tf-backend.name
  container_access_type = "private"
}
