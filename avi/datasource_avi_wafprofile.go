/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviWafProfile() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviWafProfileRead,
		Schema: map[string]*schema.Schema{
			"config": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceWafConfigSchema(),
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"files": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceWafDataFileSchema(),
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
