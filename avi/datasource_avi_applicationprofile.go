/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviApplicationProfile() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviApplicationProfileRead,
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
			"dns_service_profile": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDnsServiceApplicationProfileSchema(),
			},
			"dos_rl_profile": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDosRateLimitProfileSchema(),
			},
			"http_profile": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHTTPApplicationProfileSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"preserve_client_ip": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"preserve_client_port": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"sip_service_profile": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSipServiceApplicationProfileSchema(),
			},
			"tcp_app_profile": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceTCPApplicationProfileSchema(),
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
