provider "openstack" {
  user_name   = "admin"
  tenant_name = "admin"
  password    = "avi123"
  auth_url    = "http://10.80.3.49:5000/v3"
}

provider "avi" {
  avi_username   = "${var.avi_username}"
  avi_password   = "${var.avi_current_password}"
  avi_controller = "${var.avi_controller}"
  avi_tenant     = "admin"
}

resource "avi_useraccount" "avi_user"{
  username = "${var.avi_username}"
  old_password = "${var.avi_current_password}"
  password = "${var.avi_new_password}"
}

