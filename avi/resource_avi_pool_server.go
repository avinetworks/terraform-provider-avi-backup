/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/avinetworks/sdk/go/clients"
	"github.com/avinetworks/sdk/go/models"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func ResourceAviPoolServerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"pool_ref": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"ip": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"port": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		"type": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "V4",
		},
		"autoscaling_group_name": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"description": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"enabled": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"external_orchestration_id": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"external_uuid": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"hostname": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"location": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceGeoLocationSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"nw_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"prst_hdr_val": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"rewrite_host_header": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"vm_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

func resourceAviServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviServerCreateOrUpdate,
		Read:   ResourceAviServerRead,
		Update: resourceAviServerCreateOrUpdate,
		Delete: resourceAviServerDelete,
		Schema: ResourceAviPoolServerSchema(),
	}
}

func resourceAviServerCreateOrUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.AviClient)
	err, pUUID, poolObj, pserver := resourceAviServerReadApi(d, meta)

	// TODO: add check for err and poolObj.
	if pserver == nil {
		// not found
		newServer := models.Server{}
		if port, ok := d.GetOk("port"); ok {
			newServer.Port = int32(port.(int))
		}
		pserver = &newServer
	}
	log.Printf("[INFO] resourceAviServerCreateOrUpdate pool %v server %v", pUUID, pserver)
	// TODO set other attributes from server.
	if hostname, ok := d.GetOk("hostname"); ok {
		pserver.Hostname = hostname.(string)
	}
	if t, ok := d.GetOk("type"); ok {
		pserver.IP = &models.IPAddr{Type: t.(string), Addr: d.Get("ip").(string)}
	}
	uri := "api/pool/" + pUUID + "?include_name=true"
	var response interface{}
	patchPool := models.Pool{}
	patchPool.Name = poolObj.Name
	patchPool.TenantRef = poolObj.TenantRef
	patchPool.CloudRef = poolObj.CloudRef
	patchPool.Servers = append(patchPool.Servers, pserver)
	err = client.AviSession.Patch(uri, patchPool, "add", response)
	log.Printf("[INFO] resourceAviServerCreateOrUpdate pool %v err %v response %v", pUUID, err, response)
	if err == nil {
		err = ResourceAviServerRead(d, meta)
	}
	return err
}

func ResourceAviServerRead(d *schema.ResourceData, meta interface{}) error {
	err, pUUID, _, pserver := resourceAviServerReadApi(d, meta)
	ip := d.Get("ip").(string)
	port := d.Get("port")
	log.Printf("[INFO] pool %v ip %v port %v", pUUID, ip, port)
	if err == nil && pserver != nil {
		// TODO fix the set id to include port number. if port is not in tf then use 0
		sUUID := pUUID + ip
		d.SetId(sUUID)
		// Fill in the server information
		if pserver.Hostname != "" {
			d.Set("hostname", pserver.Hostname)
		}
		d.Set("enabled", pserver.Enabled)
		if pserver.ExternalUUID != "" {
			d.Set("external_uuid", pserver.ExternalUUID)
		}
		// Add more fields to read.
	}
	return err
}

func resourceAviServerReadApi(d *schema.ResourceData, meta interface{}) (error, string, *models.Pool, *models.Server) {
	client := meta.(*clients.AviClient)
	pUUID, pName := UUIDFromID(d.Get("pool_ref").(string))
	uri := "api/pool/" + pUUID + "?include_name=true"
	var poolObj *models.Pool
	err := client.AviSession.Get(uri, &poolObj)
	if err != nil {
		log.Printf("[ERROR] pool %v uuid %v not found", pName, pUUID)
		return err, pUUID, nil, nil
	}
	ip := d.Get("ip").(string)
	port := d.Get("port")

	// find the server in the pool object.
	var matchedServer *models.Server = nil
	for i := 0; i < len(poolObj.Servers); i++ {
		sObj := poolObj.Servers[i]
		if sObj.IP.Addr == ip {
			if (port == nil && sObj.Port == 0) || (port != nil && int32(port.(int)) == sObj.Port) {
				matchedServer = sObj
				break
			}
		}
	}

	return nil, pUUID, poolObj, matchedServer
}

func resourceAviServerDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.AviClient)
	err, pUUID, poolObj, pserver := resourceAviServerReadApi(d, meta)
	if pserver != nil {
		uri := "pool/" + pUUID + "?include_name=true&skip_default=true"
		var response interface{}
		patchPool := models.Pool{}
		patchPool.Name = poolObj.Name
		patchPool.TenantRef = poolObj.TenantRef
		patchPool.CloudRef = poolObj.CloudRef
		patchPool.Servers = append(patchPool.Servers, pserver)
		err = client.AviSession.Patch(uri, patchPool, "delete", response)
		log.Printf("[INFO] pool %v uuid %v not found", patchPool.Name, pUUID)
	}
	d.SetId("")
	return err
}
