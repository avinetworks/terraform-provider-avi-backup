/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviSecurityPolicy() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviSecurityPolicyRead,
		Schema: map[string]*schema.Schema{
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dns_attacks": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDnsAttacksSchema(),
			},
			"dns_policy_index": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_security_policy_index": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"oper_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "DETECTION"},
			"tcp_attacks": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceTcpAttacksSchema(),
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"udp_attacks": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceUdpAttacksSchema(),
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
