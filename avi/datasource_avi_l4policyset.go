/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviL4PolicySet() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviL4PolicySetRead,
		Schema: map[string]*schema.Schema{
			"created_by": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_internal_policy": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"l4_connection_policy": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceL4ConnectionPolicySchema(),
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
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
