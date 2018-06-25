package avi

import (
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func ResourceFileServiceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"http_method": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"uri": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"file_path": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
	}
}

func resourceAviFileService() *schema.Resource {
	return &schema.Resource{
		Read:   ResourceAviFileServiceRead,
		Create: ResourceAviFileServiceUploadOrDownload,
		Delete: ResourceAviFileServiceDelete,
		Schema: ResourceFileServiceSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceFileServiceImporter,
		},
	}
}

func ResourceAviFileServiceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.AviClient)
	path := "/api/fileservice"
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

func ResourceAviFileServiceUploadOrDownload(d *schema.ResourceData, meta interface{}) error {
	s := ResourceFileServiceSchema()
	err := MultipartUploadOrDownload(d, meta, s)
	if err != nil {
		log.Printf("[ERROR] ResourceAviFileServiceUploadOrDownload Error during upload/download %v\n", err)
	}
	return nil
}

func ResourceAviFileServiceDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
