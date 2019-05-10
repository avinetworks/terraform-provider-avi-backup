/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviVrfContext() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviVrfContextRead,
		Schema: map[string]*schema.Schema{
			"bgp_profile": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceBgpProfileSchema(),
			},
			"cloud_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"debugvrfcontext": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDebugVrfContextSchema(),
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"gateway_mon": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceGatewayMonitorSchema(),
			},
			"internal_gateway_monitor": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceInternalGatewayMonitorSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"static_routes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceStaticRouteSchema(),
			},
			"system_default": {
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
		},
	}
}
