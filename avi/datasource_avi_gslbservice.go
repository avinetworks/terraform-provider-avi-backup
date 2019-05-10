/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviGslbService() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviGslbServiceRead,
		Schema: map[string]*schema.Schema{
			"application_persistence_profile_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"controller_health_status_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"created_by": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"domain_names": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"down_response": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbServiceDownResponseSchema(),
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"groups": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceGslbPoolSchema(),
			},
			"health_monitor_refs": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"health_monitor_scope": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "GSLB_SERVICE_HEALTH_MONITOR_ALL_MEMBERS"},
			"hm_off": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"is_federated": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"min_members": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"num_dns_ip": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"pool_algorithm": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "GSLB_SERVICE_ALGORITHM_PRIORITY"},
			"site_persistence_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ttl": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"use_edns_client_subnet": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wildcard_match": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}
