package datadog

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestDatadogDowntime_import(t *testing.T) {
	resourceName := "datadog_downtime.foo"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDatadogDowntimeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckDatadogDowntimeConfigImported,
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

const testAccCheckDatadogDowntimeConfigImported = `
resource "datadog_downtime" "foo" {
  scope = ["host:X", "host:Y"]
  start = 1735707600
  end   = 1735765200
  timezone = "UTC"

  message = "Example Datadog downtime message."
  monitor_tags = ["*"]
}
`
