package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviFileService() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviFileServiceUploadOrDownload,
		Schema: map[string]*schema.Schema{
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
		},
	}
}
