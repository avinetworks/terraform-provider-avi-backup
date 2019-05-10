/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviAlertConfig() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviAlertConfigRead,
		Schema: map[string]*schema.Schema{
			"action_group_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"alert_rule": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAlertRuleSchema(),
			},
			"autoscale_alert": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"category": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"expiry_time": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  86400,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"obj_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"object_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"recommendation": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"rolling_window": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  300,
			},
			"source": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"summary": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"threshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"throttle": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  600,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
