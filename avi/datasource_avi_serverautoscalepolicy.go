/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviServerAutoScalePolicy() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviServerAutoScalePolicyRead,
		Schema: map[string]*schema.Schema{
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"intelligent_autoscale": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"intelligent_scalein_margin": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  40,
			},
			"intelligent_scaleout_margin": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  20,
			},
			"max_scalein_adjustment_step": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"max_scaleout_adjustment_step": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"max_size": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"min_size": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"scalein_alertconfig_refs": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"scalein_cooldown": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  300,
			},
			"scaleout_alertconfig_refs": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"scaleout_cooldown": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  300,
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"use_predicted_load": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
