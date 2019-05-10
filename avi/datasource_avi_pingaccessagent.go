/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviPingAccessAgent() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviPingAccessAgentRead,
		Schema: map[string]*schema.Schema{
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"pingaccess_pool_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"primary_server": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcePoolServerSchema(),
			},
			"properties_file_data": {
				Type:     schema.TypeString,
				Optional: true,
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
