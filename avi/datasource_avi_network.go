/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviNetwork() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviNetworkRead,
		Schema: map[string]*schema.Schema{
			"cloud_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"configured_subnets": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSubnetSchema(),
			},
			"dhcp_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"exclude_discovered_subnets": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"ip6_autocfg_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"synced_from_se": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
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
			"vcenter_dvs": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"vrf_context_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
