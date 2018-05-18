package icingamon

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/nilkonto/icingamon/iapi"
)

func TestAccCreateBasicHost(t *testing.T) {

	var testAccCreateBasicHost = fmt.Sprintf(`
		resource "icingamon_host" "tf-1" {
		hostname      = "terraform-host-1"
		address       = "10.10.10.1"
		check_command = "hostalive"
	}`)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCreateBasicHost,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHostExists("icingamon_host.tf-1"),
					testAccCheckResourceState("icingamon_host.tf-1", "hostname", "terraform-host-1"),
					testAccCheckResourceState("icingamon_host.tf-1", "address", "10.10.10.1"),
					testAccCheckResourceState("icingamon_host.tf-1", "check_command", "hostalive"),
				),
			},
		},
	})
}

func TestAccCreateVariableHost(t *testing.T) {

	var testAccCreateVariableHost = fmt.Sprintf(`
		resource "icingamon_host" "tf-3" {
		hostname = "terraform-host-3"
		address = "10.10.10.3"
		check_command = "hostalive"
		vars {
		  os = "linux"
		  osver = "1"
		  allowance = "none" }
		}`)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCreateVariableHost,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHostExists("icingamon_host.tf-3"),
					testAccCheckResourceState("icingamon_host.tf-3", "hostname", "terraform-host-3"),
					testAccCheckResourceState("icingamon_host.tf-3", "address", "10.10.10.3"),
					testAccCheckResourceState("icingamon_host.tf-3", "check_command", "hostalive"),
					testAccCheckResourceState("icingamon_host.tf-3", "vars.%", "3"),
					testAccCheckResourceState("icingamon_host.tf-3", "vars.allowance", "none"),
					testAccCheckResourceState("icingamon_host.tf-3", "vars.os", "linux"),
					testAccCheckResourceState("icingamon_host.tf-3", "vars.osver", "1"),
				),
			},
		},
	})
}

func TestAccCreateTemplateHost(t *testing.T) {
	testAccCreateTemplateHost := `resource "icingamon_host" "tf-4" {
	hostname = "terraform-host-4"
	address = "10.10.10.4"
	check_command = "hostalive"
	templates = ["generic", "az1"]
}`
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCreateTemplateHost,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHostExists("icingamon_host.tf-4"),
					testAccCheckResourceState("icingamon_host.tf-4", "hostname", "terraform-host-4"),
					testAccCheckResourceState("icingamon_host.tf-4", "address", "10.10.10.4"),
					testAccCheckResourceState("icingamon_host.tf-4", "check_command", "hostalive"),
					testAccCheckResourceState("icingamon_host.tf-4", "templates.#", "2"),
					testAccCheckResourceState("icingamon_host.tf-4", "templates.0", "generic"),
					testAccCheckResourceState("icingamon_host.tf-4", "templates.1", "az1"),
				),
			},
		},
	})
}

func testAccCheckHostExists(rn string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		resource, ok := s.RootModule().Resources[rn]
		if !ok {
			return fmt.Errorf("Host resource not found: %s", rn)
		}

		if resource.Primary.ID == "" {
			return fmt.Errorf("resource id not set")
		}

		client := testAccProvider.Meta().(*iapi.Server)
		_, err := client.GetHost(resource.Primary.ID)
		if err != nil {
			return fmt.Errorf("error getting getting host: %s", err)
		}

		return nil
	}

}
