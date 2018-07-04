package avi

import (
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func ResourceFileServiceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"uri": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"local_file": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		//upload flag to state current local file will be uploaded to remote server.
		"upload": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
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

func ResourceFileServiceImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceFileServiceSchema()
	return ResourceImporter(d, m, "fileservice", s)
}

func ResourceAviFileServiceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.AviClient)
	switch upload := d.Get("upload").(bool); upload {
	case true:
		uri := strings.Split(d.Get("uri").(string), "?")[0]
		path := "/api/fileservice?uri=controller://" + uri
		log.Printf("[DEBUG] ResourceAviFileServiceRead reading fileservice API status path %v\n", path)
		var res interface{}
		err := client.AviSession.Get(path, &res)
		log.Printf("[DEBUG]: ResourceAviFileServiceRead response: %v\n\n", res)
		if err != nil {
			log.Printf("[ERROR] ResourceAviFileServiceRead %v in GET of path %v\n", err, path)
			d.Set("upload", false)
			return err
		}
		return nil
	case false:
		local_file := d.Get("local_file").(string)
		log.Printf("[DEBUG] ResourceAviFileServiceRead reading local file %v\n", local_file)
		if _, err := os.Stat(local_file); os.IsNotExist(err) {
			log.Printf("File does not exist")
			d.Set("upload", true)
			return err
		} else {
			log.Printf("File exists")
			return nil
		}
	default:
		return nil
	}
}

func ResourceAviFileServiceCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceFileServiceSchema()
	local_file := d.Get("local_file").(string)
	err := MultipartUploadOrDownload(d, meta, s)
	if err != nil {
		log.Printf("[ERROR] ResourceAviFileServiceCreate Error during upload/download %v\n", err)
		return err
	}
	_, file := filepath.Split(local_file)
	d.SetId(file)
	return nil
}

func ResourceAviFileServiceUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] ResourceAviFileServiceUpdate")
	s := ResourceFileServiceSchema()
	err := MultipartUploadOrDownload(d, meta, s)
	if err != nil {
		log.Printf("[ERROR] ResourceAviFileServiceUpdate Error during upload/download %v\n", err)
		return err
	}
	err = ResourceAviFileServiceRead(d, meta)
	return err
}

func ResourceAviFileServiceDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.AviClient)
	local_file := d.Get("local_file").(string)
	switch upload := d.Get("upload").(bool); upload {
	case true:
		uri := strings.Split(d.Get("uri").(string), "?")[0]
		path := "/api/fileservice?uri=controller://" + uri + "/" + d.Id()
		log.Printf("[DEBUG] ResourceAviFileServiceDelete deleting file using fileservice API status path %v\n", path)
		err := client.AviSession.Delete(path)
		if err != nil {
			log.Printf("[ERROR] ResourceAviFileServiceDelete %v in Delete of path %v\n", err, path)
			return err
		}
		d.SetId("")
	case false:
		// delete file
		var err = os.Remove(local_file)
		if err != nil {
			log.Printf("[ERROR] ResourceAviFileServiceDelete Error for deleting file %v\n", local_file)
			return err
		}
		log.Printf("[INFO] ResourceAviFileServiceDelete file %v deleted\n", local_file)
		d.SetId("")
	default:
		return nil
	}
	return nil
}
