provider "azurerm" {
  subscription_id = "${var.subscription_id}"
  client_id 		  = "${var.client_id}"
  client_secret 	= "${var.client_secret}"
  tenant_id 		  = "${var.tenant_id}"
}

data "azurerm_network_interface" "controller_nic" {
  name                = "${var.project_name}-terraform-network-interface"
  resource_group_name = "${var.project_name}-terraform-resource-group"
}

data "azurerm_subnet" "terraform_subnet" {
  name                 = "${var.project_name}-terraform-subnet"
  virtual_network_name = "${var.project_name}-terraform-virtual-network"
  resource_group_name  = "${var.project_name}-terraform-resource-group"
}

provider "avi" {
  avi_username   = "${var.avi_username}"
  avi_password   = "${var.avi_password}"
  avi_controller = "${data.azurerm_network_interface.controller_nic.private_ip_address}"
  avi_tenant     = "admin"
  avi_version    = "${var.avi_version}"
}

data "avi_tenant" "default_tenant" {
  name = "admin"
}

resource "avi_cloudconnectoruser" "azure_user" {
    name       = "azureclouduser"
    tenant_ref = "${data.avi_tenant.default_tenant.id}"
    azure_userpass {
        username    = "${var.azure_username}"
        password    = "${var.azure_password}"
        tenant_name = "${var.tenant_id}"
    }
}

resource "avi_cloud" "aws_cloud_cfg" {
  name         = "AZURE"
  vtype        = "CLOUD_AZURE"
  dhcp_enabled = true
  azure_configuration {
    use_managed_disks     = true
    use_enhanced_ha       = false
    cloud_credentials_ref = "${avi_cloudconnectoruser.azure_user.id}"
    use_azure_dns         = false
    resource_group        = "${var.project_name}-terraform-resource-group"
    location              = "${var.location}"
    subscription_id       = "${var.subscription_id}"
    network_info {
        se_network_id      = "${data.azurerm_subnet.terraform_subnet.id}"
        virtual_network_id = "/subscriptions/${var.subscription_id}/resourceGroups/${var.project_name}-terraform-resource-group/providers/Microsoft.Network/virtualNetworks/${var.project_name}-terraform-virtual-network"
    }
  }
  license_tier = "ENTERPRISE_18"
  license_type = "LIC_CORES"
  tenant_ref   = "${data.avi_tenant.default_tenant.id}"
}

