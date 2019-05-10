/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviGslb() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviGslbRead,
		Schema: map[string]*schema.Schema{
			"async_interval": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"clear_on_max_retries": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  20,
			},
			"client_ip_addr_group": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbClientIpAddrGroupSchema(),
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dns_configs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDNSConfigSchema(),
			},
			"error_resync_interval": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"is_federated": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"leader_cluster_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"maintenance_mode": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"send_interval": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  15,
			},
			"send_interval_prior_to_maintenance_mode": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sites": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceGslbSiteSchema(),
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"third_party_sites": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceGslbThirdPartySiteSchema(),
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"view_id": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
		},
	}
}
