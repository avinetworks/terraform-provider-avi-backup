/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviServer() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviServerRead,
		Schema: map[string]*schema.Schema{
			"pool_ref": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ip": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "V4",
			},
			"autoscaling_group_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"availability_zone": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"discovered_networks": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDiscoveredNetworkSchema(),
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"external_orchestration_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"hostname": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"location": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGeoLocationSchema(),
			},
			"mac_address": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"nw_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"prst_hdr_val": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ratio": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"resolve_server_by_dns": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"rewrite_host_header": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"server_node": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"static": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"verify_network": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"vm_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
