/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviScheduler() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviSchedulerRead,
		Schema: map[string]*schema.Schema{
			"backup_config_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"end_date_time": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"frequency": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"frequency_unit": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"run_mode": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"run_script_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scheduler_action": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SCHEDULER_ACTION_BACKUP"},
			"start_date_time": {
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
