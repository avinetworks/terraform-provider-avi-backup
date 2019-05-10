/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviUserAccountProfile() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviUserAccountProfileRead,
		Schema: map[string]*schema.Schema{
			"account_lock_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  30,
			},
			"credentials_timeout_threshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  180,
			},
			"max_concurrent_sessions": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"max_login_failure_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  3,
			},
			"max_password_history_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  4,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
