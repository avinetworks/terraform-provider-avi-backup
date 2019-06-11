package avi

import (
	"fmt"
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"testing"
)

func TestAccDataSourceAzureRMBatchPool_complete(t *testing.T) {
	dataSourceName := "data.avi_heamthmonitor.test"

	config := testAccAVIHealthMonitorConfig

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.TestCheckResourceAttr(
					resource.TestCheckResourceAttr(dataSourceName, "name", fmt.Sprintf("testSystem-HTTP-abc")),
					resource.TestCheckResourceAttr(dataSourceName, "type", fmt.Sprintf("HEALTH_MONITOR_HTTP")),
				),
			},
		},
	})
}

const testAccAVIHealthMonitorConfig = `
data "avi_tenant" "default_tenant"{
        name= "admin"
}
resource "avi_healthmonitor" "testHealthMonitor" {
"receive_timeout" = "4"
"tenant_ref" = "${data.avi_tenant.default_tenant.id}"
"is_federated" = false
"failed_checks" = "3"
"send_interval" = "10"
"http_monitor" {
"exact_http_request" = false
"http_request" = "HEAD / HTTP/1.0"
"http_response_code" = ["HTTP_2XX","HTTP_3XX"]
}
"successful_checks" = "3"
"type" = "HEALTH_MONITOR_HTTP"
"name" = "testSystem-HTTP"
}

data "avi_healthmonitor" "test" {
  name                = "${avi_healthmonitor.test.name}"
  account_name        = "${avi_healthmonitor.test.type}"
}
`
