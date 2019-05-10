/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviApplicationPersistenceProfile() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviApplicationPersistenceProfileRead,
		Schema: map[string]*schema.Schema{
			"app_cookie_persistence_profile": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAppCookiePersistenceProfileSchema(),
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"hdr_persistence_profile": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHdrPersistenceProfileSchema(),
			},
			"http_cookie_persistence_profile": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHttpCookiePersistenceProfileSchema(),
			},
			"ip_persistence_profile": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIPPersistenceProfileSchema(),
			},
			"is_federated": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"persistence_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_hm_down_recovery": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "HM_DOWN_PICK_NEW_SERVER"},
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
