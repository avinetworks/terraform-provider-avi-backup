/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviCustomIpamDnsProfile() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviCustomIpamDnsProfileRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"script_params": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCustomParamsSchema(),
			},
			"script_uri": {
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
