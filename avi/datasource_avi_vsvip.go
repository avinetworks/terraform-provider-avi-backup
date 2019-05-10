/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviVsVip() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviVsVipRead,
		Schema: map[string]*schema.Schema{
			"cloud_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_info": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDnsInfoSchema(),
			},
			"east_west_placement": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"use_standard_alb": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vip": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVipSchema(),
			},
			"vrf_context_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vsvip_cloud_config_cksum": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
