/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviSSLProfile() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviSSLProfileRead,
		Schema: map[string]*schema.Schema{
			"accepted_ciphers": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "AES:3DES:RC4"},
			"accepted_versions": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSSLVersionSchema(),
			},
			"cipher_enums": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dhparam": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"enable_ssl_session_reuse": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"prefer_client_cipher_ordering": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"send_close_notify": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"ssl_rating": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSSLRatingSchema(),
			},
			"ssl_session_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  86400,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceTagSchema(),
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SSL_PROFILE_TYPE_APPLICATION"},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
