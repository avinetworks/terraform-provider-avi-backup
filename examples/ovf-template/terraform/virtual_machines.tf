// example_virtual_machines creates a single virtual machine on each individual
// host.
resource "vsphere_virtual_machine" "example_virtual_machines" {
  count            = "${length(var.esxi_hosts)}"
  name             = "${var.virtual_machine_name_prefix}${count.index}"
  resource_pool_id = "${data.vsphere_resource_pool.example_resource_pool.id}"
  host_system_id   = "${data.vsphere_host.example_hosts.*.id[count.index]}"
  datastore_id     = "${data.vsphere_datastore.example_datastore.id}"
  num_cpus = 2
  memory   = 1024
  guest_id = "${data.vsphere_virtual_machine.example_template.guest_id}"
  wait_for_guest_net_timeout = "0"
  force_power_off = "false"
  scsi_type = "${data.vsphere_virtual_machine.example_template.scsi_type}"
  network_interface {
    network_id   = "${data.vsphere_network.example_network.id}"
    adapter_type = "${data.vsphere_virtual_machine.example_template.network_interface_types[0]}"
  }

  disk {
    label = "disk0"
    size  = "${data.vsphere_virtual_machine.example_template.disks.0.size}"
    eagerly_scrub    = "${data.vsphere_virtual_machine.example_template.disks.0.eagerly_scrub}"
    thin_provisioned = "${data.vsphere_virtual_machine.example_template.disks.0.thin_provisioned}"
  }

  clone {
    template_uuid = "${data.vsphere_virtual_machine.example_template.id}"
  }
  vapp {
      properties {
        "mgmt-ip" = "10.10.28.132"
        "mgmt-mask" = "23"
        "default-gw" = "10.10.28.1"
      }
  }
}
