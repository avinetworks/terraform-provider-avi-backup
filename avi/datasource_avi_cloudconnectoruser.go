/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviCloudConnectorUser() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviCloudConnectorUserRead,
		Schema: map[string]*schema.Schema{
			"azure_serviceprincipal": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAzureServicePrincipalCredentialsSchema(),
			},
			"azure_userpass": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAzureUserPassCredentialsSchema(),
			},
			"gcp_credentials": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGCPCredentialsSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"oci_credentials": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceOCICredentialsSchema(),
			},
			"password": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"private_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"public_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tencent_credentials": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceTencentCredentialsSchema(),
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
