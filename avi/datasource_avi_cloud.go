/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviCloud() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviCloudRead,
		Schema: map[string]*schema.Schema{
			"apic_configuration": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAPICConfigurationSchema(),
			},
			"apic_mode": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"autoscale_polling_interval": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60,
			},
			"aws_configuration": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAwsConfigurationSchema(),
			},
			"azure_configuration": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAzureConfigurationSchema(),
			},
			"cloudstack_configuration": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCloudStackConfigurationSchema(),
			},
			"custom_tags": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCustomTagSchema(),
			},
			"dhcp_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"dns_provider_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"docker_configuration": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDockerConfigurationSchema(),
			},
			"east_west_dns_provider_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"east_west_ipam_provider_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enable_vip_static_routes": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"gcp_configuration": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGCPConfigurationSchema(),
			},
			"ip6_autocfg_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"ipam_provider_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"license_tier": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"license_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"linuxserver_configuration": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceLinuxServerConfigurationSchema(),
			},
			"mtu": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1500,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"nsx_configuration": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceNsxConfigurationSchema(),
			},
			"obj_name_prefix": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"openstack_configuration": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceOpenStackConfigurationSchema(),
			},
			"oshiftk8s_configuration": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceOShiftK8SConfigurationSchema(),
			},
			"prefer_static_routes": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"proxy_configuration": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceProxyConfigurationSchema(),
			},
			"rancher_configuration": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRancherConfigurationSchema(),
			},
			"se_group_template_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"state_based_dns_registration": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vca_configuration": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcevCloudAirConfigurationSchema(),
			},
			"vcenter_configuration": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcevCenterConfigurationSchema(),
			},
			"vtype": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
