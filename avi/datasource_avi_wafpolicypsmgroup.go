/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviWafPolicyPSMGroup() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviWafPolicyPSMGroupRead,
		Schema: map[string]*schema.Schema{
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"enable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"hit_action": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "WAF_ACTION_NO_OP"},
			"is_learning_group": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"locations": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceWafPSMLocationSchema(),
			},
			"miss_action": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "WAF_ACTION_NO_OP"},
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
