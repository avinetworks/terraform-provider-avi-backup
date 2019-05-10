/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviPKIProfile() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviPKIProfileRead,
		Schema: map[string]*schema.Schema{
			"ca_certs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSSLCertificateSchema(),
			},
			"created_by": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"crl_check": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"crls": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCRLSchema(),
			},
			"ignore_peer_chain": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
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
			"validate_only_leaf_crl": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
		},
	}
}
