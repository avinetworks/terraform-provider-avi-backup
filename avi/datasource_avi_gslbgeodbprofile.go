/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviGslbGeoDbProfile() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviGslbGeoDbProfileRead,
		Schema: map[string]*schema.Schema{
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"entries": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceGslbGeoDbEntrySchema(),
			},
			"is_federated": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
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
