# terraform {
#   backend "azurerm" {
#     resource_group_name  = "dev-eastus-rsg"
#     storage_account_name = "deveastusstorage"
#     container_name       = "deveastuscontainer"
#     key                  = "dev.terraform.tfstate"
#   }
# }