resource "azurerm_resource_group" "myrsg" {
  name     = format("%s-%s-%s", terraform.workspace, var.location, "rsg")
  location = var.location
  tags     = local.required_tags
}
variable "location" {}

variable "tags" {
  type        = map(any)
  description = "adding tags for resources by using local function"
}

resource "azurerm_storage_account" "mystorage" {
  name                     = lower(format("%s%s%s", terraform.workspace, var.location, "storage"))
  resource_group_name      = azurerm_resource_group.myrsg.name
  location                 = azurerm_resource_group.myrsg.location
  account_tier             = "Standard"
  account_replication_type = "LRS"
  tags                     = local.required_tags
  depends_on = [
    azurerm_resource_group.myrsg
  ]
}

resource "azurerm_storage_container" "container" {
  name                  = lower(format("%s%s%s", terraform.workspace, var.location, "container"))
  storage_account_name  = azurerm_storage_account.mystorage.name
  container_access_type = "private"
  #tags = local.required_tags ==> tags are not expected here
  depends_on = [
    azurerm_resource_group.myrsg,
    azurerm_storage_account.mystorage
  ]

}

resource "azurerm_storage_blob" "blob" {
  name                   = lower(format("%s%s%s", terraform.workspace, var.location, "blob"))
  storage_container_name = azurerm_storage_container.container.name
  storage_account_name   = azurerm_storage_account.mystorage.name
  type                   = "Block"
  #tags = local.required_tags ==> tags are not expected here
  depends_on = [
    azurerm_resource_group.myrsg,
    azurerm_storage_account.mystorage
  ]


}
#{"Account" = "Storage","Subscription" = "Dev","Application" = "web"}

locals {
  required_tags = {
    Account      = var.tags["Account"]
    Subscription = var.tags["Subscription"]
    Application  = var.tags["Application"]
  }
}

