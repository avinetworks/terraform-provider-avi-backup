package avi

import (
	"github.com/avinetworks/sdk/go/clients"
	//"github.com/avinetworks/sdk/go/models"
	"github.com/avinetworks/sdk/go/models"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func resourceAviPool() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviPoolCreate,
		Read:   resourceAviPoolRead,
		//Update: resourceAviPoolUpdate,
		Delete: resourceAviPoolDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "Name of Avi Pool",
				Required:    true,
				ForceNew:    true,
			},
			"uuid": {
				Type:        schema.TypeString,
				Description: "UUID of Avi Pool",
				Computed:    true,
			},
		},
	}
}

func resourceAviPoolRead(d *schema.ResourceData, meta interface{}) error {
	log.Println("[INFO] resourceAviPoolRead Avi Client")
	client := meta.(*clients.AviClient)
	uuid := d.Id()
	obj, err := client.Pool.Get(uuid)
	if err != nil {
		return err
	}
	d.Set("name", obj.Name)
	d.Set("uuid", obj.UUID)
	return nil
}

func resourceAviPoolUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceAviPoolRead(d, meta)
}

func resourceAviPoolCreate(d *schema.ResourceData, meta interface{}) error {
	log.Println("[INFO] resourceAviPoolCreate: Avi Client")
	client := meta.(*clients.AviClient)
	obj := models.Pool{}
	obj.Name = d.Get("name").(string)
	log.Println("[INFO] resourceAviPoolCreate: calling create api")
	new_obj, err := client.Pool.Create(&obj)
	if err != nil {
		log.Println("[INFO] resourceAviPoolCreate: Error")
		return err
	}
	d.SetId(new_obj.UUID)
	d.Set("uuid", new_obj.UUID)
	log.Println("[INFO] resourceAviPoolCreate: Create success")
	return nil
}

func resourceAviPoolDelete(d *schema.ResourceData, meta interface{}) error {
	log.Println("[INFO] resourceAviPoolRead Avi Client")
	client := meta.(*clients.AviClient)
	uuid := d.Id()
	err := client.Pool.Delete(uuid)
	if err != nil {
		log.Println("[INFO] resourceAviPoolDelete not found")
	}
	d.SetId("")
	return nil
}
