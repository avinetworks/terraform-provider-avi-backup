/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviHealthMonitor() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviHealthMonitorRead,
		Schema: map[string]*schema.Schema{
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dns_monitor": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHealthMonitorDNSSchema(),
			},
			"external_monitor": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHealthMonitorExternalSchema(),
			},
			"failed_checks": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  2,
			},
			"http_monitor": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHealthMonitorHttpSchema(),
			},
			"https_monitor": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHealthMonitorHttpSchema(),
			},
			"is_federated": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"monitor_port": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"receive_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  4,
			},
			"send_interval": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"sip_monitor": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHealthMonitorSIPSchema(),
			},
			"successful_checks": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  2,
			},
			"tcp_monitor": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHealthMonitorTcpSchema(),
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
			"udp_monitor": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHealthMonitorUdpSchema(),
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
