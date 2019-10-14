output "tf_storage_account_access_key" {
  value = azurerm_storage_account.tf-backend.primary_access_key
}

output "tf_storage_container_name" {
  value = azurerm_storage_container.tf-storage-container.name
}
