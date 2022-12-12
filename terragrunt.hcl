// remote_state {
//     backend = "azurerm"
//     config = {
//         key = "${path_relative_to_include()}/terraform.tfstate"
//         resource_group_name  = "dev-eastus-rsg"
//         storage_account_name = "deveastusstorage"
//         container_name       = "deveastuscontainer"
//     }
//     generate = {
//         path      = "backend.tf"
//         if_exists = "overwrite_terragrunt"
//     }
//     dependency "vpc" {
//         config_path = "../terraform"
// }
    
// }