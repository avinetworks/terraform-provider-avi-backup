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

func ResourceServiceEnginePolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"se_group_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"tenant_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"vrf_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviServiceEnginePolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviServiceEnginePolicyCreate,
		Read:   ResourceAviServiceEnginePolicyRead,
		Update: resourceAviServiceEnginePolicyUpdate,
		Delete: resourceAviServiceEnginePolicyDelete,
		Schema: ResourceServiceEnginePolicySchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceServiceEnginePolicyImporter,
		},
	}
}

func ResourceServiceEnginePolicyImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceServiceEnginePolicySchema()
	return ResourceImporter(d, m, "serviceenginepolicy", s)
}

func ResourceAviServiceEnginePolicyRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceServiceEnginePolicySchema()
	err := ApiRead(d, meta, "serviceenginepolicy", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviServiceEnginePolicyCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceServiceEnginePolicySchema()
	err := ApiCreateOrUpdate(d, meta, "serviceenginepolicy", s)
	if err == nil {
		err = ResourceAviServiceEnginePolicyRead(d, meta)
	}
	return err
}

func resourceAviServiceEnginePolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceServiceEnginePolicySchema()
	var err error
	err = ApiCreateOrUpdate(d, meta, "serviceenginepolicy", s)
	if err == nil {
		err = ResourceAviServiceEnginePolicyRead(d, meta)
	}
	return err
}

func resourceAviServiceEnginePolicyDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "serviceenginepolicy"
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviServiceEnginePolicyDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
