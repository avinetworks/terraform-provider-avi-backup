package avi

import (
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strings"
)

func ResourceFileServiceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"http_method": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"uri": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"file_path": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		//Force mode to overwrite the remote file if it exists on server.
		"force": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default: false,
		},
	}
}

func resourceAviFileService() *schema.Resource {
	return &schema.Resource{
		Read:   ResourceAviFileServiceRead,
		Create: ResourceAviFileServiceCreate,
		Update: ResourceAviFileServiceUpdate,
		Delete: ResourceAviFileServiceDelete,
		Schema: ResourceFileServiceSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceFileServiceImporter,
		},
	}
}

func ResourceAviFileServiceUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] ResourceAviFileServiceUpdate")
	s := ResourceFileServiceSchema()
	switch force := d.Get("force").(bool); force {
	case true:
		err := MultipartUploadOrDownload(d, meta, s)
		if err != nil {
			log.Printf("[ERROR] ResourceAviFileServiceUpdate Error during upload/download %v\n", err)
		}
		return nil
	case false:
		log.Printf("[INFO] ResourceAviFileServiceUpdate File already exists on server\n")
		return nil
	default:
		err := MultipartUploadOrDownload(d, meta, s)
		if err != nil {
			log.Printf("[ERROR] ResourceAviFileServiceUpdate Error during upload/download %v\n", err)
		}
		return nil
	}
}

func ResourceAviFileServiceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.AviClient)
	uri := strings.Split(d.Get("uri").(string), "?")[0]
	path := "/api/fileservice?uri=controller://" + uri
	log.Printf("[DEBUG] ResourceAviFileServiceRead reading fileservice API status path %v\n", path)
	var res interface{}
	err := client.AviSession.Get(path, &res)
	log.Printf("[DEBUG]: ResourceAviFileServiceRead response: %v\n\n", res)
	if err != nil {
		log.Printf("[ERROR] ResourceAviFileServiceRead %v in GET of path %v\n", err, path)
		return err
	}
	return nil
}

func ResourceFileServiceImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceFileServiceSchema()
	return ResourceImporter(d, m, "fileservice", s)
}

func ResourceAviFileServiceCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceFileServiceSchema()
	err := MultipartUploadOrDownload(d, meta, s)
	if err != nil {
		log.Printf("[ERROR] ResourceAviFileServiceCreate Error during upload/download %v\n", err)
	}
	return nil
}

func ResourceAviFileServiceDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.AviClient)
	uri := strings.Split(d.Get("uri").(string), "?")[0]
	path := "/api/fileservice?uri=controller://" + uri + "/" + d.Id()
	log.Printf("[DEBUG] ResourceAviFileServiceDelete deleting file using fileservice API status path %v\n", path)
	err := client.AviSession.Delete(path)
	if err != nil {
		log.Printf("[ERROR] ResourceAviFileServiceDelete %v in Delete of path %v\n", err, path)
		return err
	}
	d.SetId("")
	return nil
}
