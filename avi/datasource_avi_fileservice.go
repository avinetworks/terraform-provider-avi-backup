package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviFileService() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviFileServiceCreate,
		Schema: map[string]*schema.Schema{
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
		},
	}
}
