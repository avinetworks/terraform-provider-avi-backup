/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviHTTPPolicySet() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviHTTPPolicySetRead,
		Schema: map[string]*schema.Schema{
			"cloud_config_cksum": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"created_by": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_request_policy": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHTTPRequestPolicySchema(),
			},
			"http_response_policy": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHTTPResponsePolicySchema(),
			},
			"http_security_policy": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHTTPSecurityPolicySchema(),
			},
			"is_internal_policy": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
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
