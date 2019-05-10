/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviServiceEngine() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviServiceEngineRead,
		Schema: map[string]*schema.Schema{
			"availability_zone": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cloud_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"container_mode": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"container_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "CONTAINER_TYPE_HOST"},
			"controller_created": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"controller_ip": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"data_vnics": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourcevNICSchema(),
			},
			"enable_state": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SE_STATE_ENABLED"},
			"flavor": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"host_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hypervisor": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"mgmt_vnic": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcevNICSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "VM name unknown"},
			"resources": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeResourcesSchema(),
			},
			"se_group_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
