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

func ResourceWafCRSSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"description": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"groups": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceWafRuleGroupSchema(),
		},
		"integrity": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"release_date": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
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
		"version": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

func resourceAviWafCRS() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviWafCRSCreate,
		Read:   ResourceAviWafCRSRead,
		Update: resourceAviWafCRSUpdate,
		Delete: resourceAviWafCRSDelete,
		Schema: ResourceWafCRSSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceWafCRSImporter,
		},
	}
}

func ResourceWafCRSImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceWafCRSSchema()
	return ResourceImporter(d, m, "wafcrs", s)
}

func ResourceAviWafCRSRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWafCRSSchema()
	err := ApiRead(d, meta, "wafcrs", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviWafCRSCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWafCRSSchema()
	err := ApiCreateOrUpdate(d, meta, "wafcrs", s)
	if err == nil {
		err = ResourceAviWafCRSRead(d, meta)
	}
	return err
}

func resourceAviWafCRSUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWafCRSSchema()
	var err error
	err = ApiCreateOrUpdate(d, meta, "wafcrs", s)
	if err == nil {
		err = ResourceAviWafCRSRead(d, meta)
	}
	return err
}

func resourceAviWafCRSDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "wafcrs"
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviWafCRSDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
