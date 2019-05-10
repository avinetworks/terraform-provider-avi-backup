/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviErrorPageBody() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviErrorPageBodyRead,
		Schema: map[string]*schema.Schema{
			"error_page_body": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"format": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "ERROR_PAGE_FORMAT_HTML"},
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
