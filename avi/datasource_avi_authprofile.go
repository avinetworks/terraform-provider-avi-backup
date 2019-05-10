/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviAuthProfile() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviAuthProfileRead,
		Schema: map[string]*schema.Schema{
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"http": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAuthProfileHTTPClientParamsSchema(),
			},
			"ldap": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceLdapAuthSettingsSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"pa_agent_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"saml": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSamlSettingsSchema(),
			},
			"tacacs_plus": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceTacacsPlusAuthSettingsSchema(),
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": {
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
