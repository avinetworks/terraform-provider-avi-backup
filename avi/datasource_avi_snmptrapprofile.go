/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviSnmpTrapProfile() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviSnmpTrapProfileRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"trap_servers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceSnmpTrapServerSchema(),
			},
			"uuid": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
