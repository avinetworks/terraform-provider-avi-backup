package avi

import (
	"github.com/avinetworks/sdk/go/clients"
	"github.com/avinetworks/sdk/go/models"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	//"reflect"
)

func resourceAviPool() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviPoolCreate,
		Read:   resourceAviPoolRead,
		Update: resourceAviPoolUpdate,
		Delete: resourceAviPoolDelete,
		Schema: map[string]*schema.Schema{
			"obj": {
				//Type:        schema.TypeSet,
				Type:        schema.TypeMap,
				Description: "Pool Object",
				Optional:    true,
				//Set: func(v interface{}) int {
				//	return 0
				//},
				// Desired Avi Object
				//Elem: &models.Pool{},
				/*
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"name": &schema.Schema{
								Type:     schema.TypeString,
								Required: true,
							},
							"health_monitor_refs": &schema.Schema{
								Type:     schema.TypeList,
								Optional: true,
								Elem:     &schema.Schema{Type: schema.TypeString},
							},
							"servers": &schema.Schema{
								Type:     schema.TypeList,
								Optional: true,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"ip": &schema.Schema{
											Type:     schema.TypeSet,
											Optional: true,
											Elem: &schema.Resource{
												Schema: map[string]*schema.Schema{
													"type": &schema.Schema{
														Type:     schema.TypeString,
														Optional: true,
													},
													"addr": &schema.Schema{
														Type:     schema.TypeString,
														Optional: true,
													},
												},
											},
										},
										"port": &schema.Schema{
											Type:     schema.TypeInt,
											Optional: true,
										},
									},
								},
							},
						},
					},*/
			},
		},
	}
}

func resourceAviPoolRead(d *schema.ResourceData, meta interface{}) error {
	log.Println("[INFO] resourceAviPoolRead Avi Client")
	client := meta.(*clients.AviClient)
	_, err := client.Pool.Get(d.Id())
	if err != nil {
		d.SetId("")
		return nil
	}
	// no need to set the ID
	log.Println("[INFO] resourceAviPoolRead Updating obj")
	//d.Set("obj", obj)
	return nil
}

func resourceAviPoolUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
	/*
		log.Println("[INFO] resourceAviPoolUpdate: Avi Client")
		client := meta.(*clients.AviClient)
		if aobj, ok := d.GetOk("obj"); ok {
			interfaceSlice := aobj.(*schema.Set).List()
			obj := interfaceSlice[0]
			var robj *models.Pool
			api_path := "api/pool"
			log.Printf("[DEBUG] resourceAviPoolUpdate: data %v\n", obj)
			err := client.AviSession.Put(api_path, obj, &robj)
			if err != nil {
				log.Println("[INFO] resourceAviPoolUpdate: Error")
				return err
			}
			log.Println("[INFO] resourceAviPoolUpdate: Create success")
			d.Set("obj", robj)
		}
		return nil
	*/
}

func resourceAviPoolCreate(d *schema.ResourceData, meta interface{}) error {
	log.Println("[INFO] resourceAviPoolCreate: Avi Client")
	client := meta.(*clients.AviClient)
	var robj *models.Pool
	api_path := "api/pool"

	aobj := d.Get("obj")
	log.Printf("[DEBUG] resourceAviPoolCreate: data %v\n", aobj)
	err := client.AviSession.Post(api_path, aobj, &robj)
	if err != nil {
		log.Println("[INFO] resourceAviPoolCreate: Error")
		return err
	}
	d.SetId(robj.UUID)
	log.Println("[INFO] resourceAviPoolCreate: Create success")

	/*
		if aobj, ok := d.GetOk("obj"); ok {
			log.Printf("[DEBUG] resourceAviPoolCreate: data %v %v\n", aobj, reflect.TypeOf(aobj))
			interfaceSlice := aobj.(*schema.Set).List()
			log.Printf("[DEBUG] resourceAviPoolCreate: sliced data %v\n", interfaceSlice)
			obj := interfaceSlice[0]
			log.Printf("[DEBUG] resourceAviPoolCreate: POST data %v %v\n", obj, reflect.TypeOf(obj))
			err := client.AviSession.Post(api_path, aobj, &robj)
			log.Printf("[DEBUG] resourceAviPoolCreate: POST err %v %v\n", err, robj)
			if err != nil {
				log.Println("[ERROR] resourceAviPoolCreate: Error")
				return err
			}
			d.SetId(robj.UUID)
			log.Println("[INFO] resourceAviPoolCreate: Create success")
			return nil
		}
	*/
	return nil
}

func resourceAviPoolDelete(d *schema.ResourceData, meta interface{}) error {
	log.Println("[INFO] resourceAviPoolRead Avi Client")
	client := meta.(*clients.AviClient)
	uuid := d.Id()

	if uuid != "" {
		err := client.Pool.Delete(uuid)
		if err != nil {
			log.Println("[INFO] resourceAviPoolDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
