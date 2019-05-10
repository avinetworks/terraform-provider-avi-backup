/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviTenant() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviTenantRead,
		Schema: map[string]*schema.Schema{
			"config_settings": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceTenantConfigurationSchema(),
			},
			"created_by": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"local": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
