/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviPoolGroup() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviPoolGroupRead,
		Schema: map[string]*schema.Schema{
			"cloud_config_cksum": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cloud_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"created_by": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"deployment_policy_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"fail_action": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceFailActionSchema(),
			},
			"implicit_priority_labels": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"members": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourcePoolGroupMemberSchema(),
			},
			"min_servers": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"priority_labels_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"service_metadata": {
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
