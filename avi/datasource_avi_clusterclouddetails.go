/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviClusterCloudDetails() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviClusterCloudDetailsRead,
		Schema: map[string]*schema.Schema{
			"azure_info": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceAzureClusterInfoSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
