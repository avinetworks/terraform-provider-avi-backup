/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviIpAddrGroup() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviIpAddrGroupRead,
		Schema: map[string]*schema.Schema{
			"addrs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"apic_epg_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"country_codes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip_ports": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrPortSchema(),
			},
			"marathon_app_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"marathon_service_port": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"prefixes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
			"ranges": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrRangeSchema(),
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
		},
	}
}
