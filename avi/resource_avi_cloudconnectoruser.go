/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strings"
)

func ResourceCloudConnectorUserSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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
			Required: true,
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
	}
}

func resourceAviCloudConnectorUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviCloudConnectorUserCreate,
		Read:   ResourceAviCloudConnectorUserRead,
		Update: resourceAviCloudConnectorUserUpdate,
		Delete: resourceAviCloudConnectorUserDelete,
		Schema: ResourceCloudConnectorUserSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceCloudConnectorUserImporter,
		},
	}
}

func ResourceCloudConnectorUserImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceCloudConnectorUserSchema()
	return ResourceImporter(d, m, "cloudconnectoruser", s)
}

func ResourceAviCloudConnectorUserRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCloudConnectorUserSchema()
	err := ApiRead(d, meta, "cloudconnectoruser", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviCloudConnectorUserCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCloudConnectorUserSchema()
	err := ApiCreateOrUpdate(d, meta, "cloudconnectoruser", s)
	if err == nil {
		err = ResourceAviCloudConnectorUserRead(d, meta)
	}
	return err
}

func resourceAviCloudConnectorUserUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCloudConnectorUserSchema()
	var err error
	err = ApiCreateOrUpdate(d, meta, "cloudconnectoruser", s)
	if err == nil {
		err = ResourceAviCloudConnectorUserRead(d, meta)
	}
	return err
}

func resourceAviCloudConnectorUserDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "cloudconnectoruser"
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviCloudConnectorUserDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
