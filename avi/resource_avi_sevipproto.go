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

func ResourceSeVipProtoSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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
			Computed: true,
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
	}
}

func resourceAviSeVipProto() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviSeVipProtoCreate,
		Read:   ResourceAviSeVipProtoRead,
		Update: resourceAviSeVipProtoUpdate,
		Delete: resourceAviSeVipProtoDelete,
		Schema: ResourceSeVipProtoSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceSeVipProtoImporter,
		},
	}
}

func ResourceSeVipProtoImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceSeVipProtoSchema()
	return ResourceImporter(d, m, "sevipproto", s)
}

func ResourceAviSeVipProtoRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSeVipProtoSchema()
	err := ApiRead(d, meta, "sevipproto", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviSeVipProtoCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSeVipProtoSchema()
	err := ApiCreateOrUpdate(d, meta, "sevipproto", s)
	if err == nil {
		err = ResourceAviSeVipProtoRead(d, meta)
	}
	return err
}

func resourceAviSeVipProtoUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSeVipProtoSchema()
	var err error
	err = ApiCreateOrUpdate(d, meta, "sevipproto", s)
	if err == nil {
		err = ResourceAviSeVipProtoRead(d, meta)
	}
	return err
}

func resourceAviSeVipProtoDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "sevipproto"
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviSeVipProtoDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
