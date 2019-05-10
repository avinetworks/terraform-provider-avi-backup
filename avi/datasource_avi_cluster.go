/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviCluster() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviClusterRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"nodes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceClusterNodeSchema(),
			},
			"rejoin_nodes_automatically": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
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
			"virtual_ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
			},
		},
	}
}
