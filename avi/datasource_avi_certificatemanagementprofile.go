/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviCertificateManagementProfile() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviCertificateManagementProfileRead,
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
			"script_path": {
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
