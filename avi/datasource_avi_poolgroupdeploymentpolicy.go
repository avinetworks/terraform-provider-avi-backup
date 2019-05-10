/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviPoolGroupDeploymentPolicy() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviPoolGroupDeploymentPolicyRead,
		Schema: map[string]*schema.Schema{
			"auto_disable_old_prod_pools": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"evaluation_duration": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  300,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"rules": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourcePGDeploymentRuleSchema(),
			},
			"scheme": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "BLUE_GREEN"},
			"target_test_traffic_ratio": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"test_traffic_ratio_rampup": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"webhook_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
