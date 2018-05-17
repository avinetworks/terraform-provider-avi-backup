/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviSeVipProto() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviSeVipProtoRead,
		Schema: map[string]*schema.Schema{
			"con_info": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceConVipInfoSchema(),
			},
			"num_consumers_sharing_vip": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"primary_se_info": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSeVipInfoSchema(),
			},
			"secondary_se_info": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSeVipInfoSchema(),
			},
			"standby_se_info": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSeVipInfoSchema(),
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vip6": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}
