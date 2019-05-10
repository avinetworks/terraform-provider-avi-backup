/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviWafPolicy() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviWafPolicyRead,
		Schema: map[string]*schema.Schema{
			"allow_mode_delegation": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"created_by": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"crs_groups": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceWafRuleGroupSchema(),
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"enable_app_learning": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"failure_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "WAF_FAILURE_MODE_OPEN"},
			"mode": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "WAF_MODE_DETECTION_ONLY"},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"paranoia_level": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "WAF_PARANOIA_LEVEL_LOW"},
			"positive_security_model": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceWafPositiveSecurityModelSchema(),
			},
			"post_crs_groups": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceWafRuleGroupSchema(),
			},
			"pre_crs_groups": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceWafRuleGroupSchema(),
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
			"waf_crs_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"waf_profile_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"whitelist": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceWafPolicyWhitelistSchema(),
			},
		},
	}
}
