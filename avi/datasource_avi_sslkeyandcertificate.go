/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviSSLKeyAndCertificate() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviSSLKeyAndCertificateRead,
		Schema: map[string]*schema.Schema{
			"ca_certs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCertificateAuthoritySchema(),
			},
			"certificate": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSSLCertificateSchema(),
			},
			"certificate_base64": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"certificate_management_profile_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"created_by": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dynamic_params": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCustomParamsSchema(),
			},
			"enckey_base64": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"enckey_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"format": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SSL_PEM"},
			"hardwaresecuritymodulegroup_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"key_base64": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"key_params": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSSLKeyParamsSchema(),
			},
			"key_passphrase": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SSL_CERTIFICATE_FINISHED"},
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
