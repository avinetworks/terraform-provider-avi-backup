# Configure vSphere provider
provider "vsphere" {
  vsphere_server = "${var.vsphere_vcenter}"
  user = "${var.vsphere_user}"
  password = "${var.vsphere_password}"

  allow_unverified_ssl = "${var.vsphere_unverified_ssl}"
}

data "vsphere_datacenter" "dc" {
  name = "${var.vsphere_datacenter}"
}

data "vsphere_datastore" "datastore" {
  name          = "${var.vsphere_datastore}"
  datacenter_id = "${data.vsphere_datacenter.dc.id}"
}

data "vsphere_resource_pool" "pool" {
  name          = "${var.vsphere_resource_pool}"
  datacenter_id = "${data.vsphere_datacenter.dc.id}"
}

data "vsphere_network" "network" {
  name          = "${var.vsphere_mgmt_group}"
  datacenter_id = "${data.vsphere_datacenter.dc.id}"
}

data "vsphere_host" "host" {
  name          = "${var.vsphere_specific_host}"
  datacenter_id = "${data.vsphere_datacenter.dc.id}"
}
#data "vsphere_virtual_machine" "template" {
#  name          = "${var.vsphere_destination_file}"
#  datacenter_id = "${data.vsphere_datacenter.dc.id}"
#}


resource "vsphere_file" "ubuntu_disk_upload" {
  datacenter       = "${var.vsphere_datacenter}"
  datastore        = "${var.vsphere_datastore}"
  source_file      = "${var.vsphere_source_file}"
  destination_file = "${var.vsphere_destination_file}"
  create_directories = "True"
}


# Create a vSphere VM in the terraform folder
resource "vsphere_virtual_machine" "terraform-ubuntu" {
  # VM placement
//  host = "${data.vsphere_host.host.id}"
  name = "${var.vsphere_vm_name}"
  resource_pool_id = "${data.vsphere_resource_pool.pool.id}"
  datastore_id     = "${data.vsphere_datastore.datastore.id}"

  # VM resources
  num_cpus = "${var.vsphere_num_cpus}"
  memory = "${var.vsphere_memory_size}"

  # VM storage
  disk {
    label             = "disk1"
    size              = "20"
  }

  # VM networking
  network_interface {
    network_id   = "${data.vsphere_network.network.id}"
  }

}
//resource "vsphere_virtual_machine" "vm" {
//  name             = "terraform-test"
//  resource_pool_id = "${data.vsphere_resource_pool.pool.id}"
//  datastore_id     = "${data.vsphere_datastore.datastore.id}"
//
//  num_cpus = 2
//  memory   = 1024
//  guest_id = "${data.vsphere_virtual_machine.template_from_ovf.guest_id}"
//
//  scsi_type = "${data.vsphere_virtual_machine.template_from_ovf.scsi_type}"
//
//  network_interface {
//    network_id   = "${data.vsphere_network.network.id}"
//    adapter_type = "${data.vsphere_virtual_machine.template_from_ovf.network_interface_types[0]}"
//  }
//
//  disk {
//    label             = "disk0"
//    size             = "${data.vsphere_virtual_machine.template_from_ovf.disks.0.size}"
//  }
//
//  clone {
//    template_uuid = "${data.vsphere_virtual_machine.template_from_ovf.id}"
//  }
//
//  vapp {
//    properties {
//      "guestinfo.hostname"                        = "terraform-test.foobar.local"
//      "guestinfo.interface.0.label"                = "${var.vsphere_port_group}"
//      "guestinfo.interface.0.ip.0.address"        = "${var.vsphere_ipv4_address}"
//      "guestinfo.interface.0.route.0.gateway"     = "${var.vsphere_ipv4_gateway}"
//      "guestinfo.interface.0.route.0.destination" = ""
//      "guestinfo.dns.server.0"                    = "${var.vsphere_dns_servers}"
//    }
//  }
//}