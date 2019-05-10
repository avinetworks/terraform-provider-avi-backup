/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviBackupConfiguration() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviBackupConfigurationRead,
		Schema: map[string]*schema.Schema{
			"aws_access_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"aws_bucket_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"aws_secret_access": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"backup_file_prefix": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"backup_passphrase": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"maximum_backups_stored": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  4,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"remote_directory": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"remote_hostname": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"save_local": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"ssh_user_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"upload_to_remote_host": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"upload_to_s3": {
				Type:     schema.TypeBool,
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
