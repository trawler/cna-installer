provider "azurerm" {
  version = "1.35.0"
}

resource "azurerm_resource_group" "cna-backend" {
  name     = var.k8s_resource_group_name
  location = var.az_location
}


// Azure terraform backend storage account
resource "azurerm_storage_account" "tf-cna-backend" {
  name                     = format("%stfstorage", var.cluster_owner)
  resource_group_name      = azurerm_resource_group.cna-backend.name
  location                 = azurerm_resource_group.cna-backend.location
  account_tier             = "Standard"
  account_replication_type = "LRS"

  tags = {
    environment = "dev"
  }
}

resource "azurerm_storage_container" "tf-storage-container" {
  name                  = "cna-tfstate"
  storage_account_name  = azurerm_storage_account.tf-cna-backend.name
  container_access_type = "private"
}
