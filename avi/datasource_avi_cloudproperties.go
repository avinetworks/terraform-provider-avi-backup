/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviCloudProperties() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviCloudPropertiesRead,
		Schema: map[string]*schema.Schema{
			"cc_props": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCC_PropertiesSchema(),
			},
			"cc_vtypes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"hyp_props": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceHypervisor_PropertiesSchema(),
			},
			"info": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCloudInfoSchema(),
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
