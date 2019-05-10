/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviAutoScaleLaunchConfig() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviAutoScaleLaunchConfigRead,
		Schema: map[string]*schema.Schema{
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"image_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"mesos": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAutoScaleMesosSettingsSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"openstack": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAutoScaleOpenStackSettingsSchema(),
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"use_external_asg": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
