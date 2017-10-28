package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/rebuy-de/terraform-provider-graylog/pkg/graylog"
	"github.com/rebuy-de/terraform-provider-graylog/pkg/testutils"
)

func TestAccInput(t *testing.T) {
	var name string
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccInputCheckDestroy(name),
		Steps: []resource.TestStep{
			{
				Config: testAccInputConfig_gelfUDP,
				Check: resource.ComposeTestCheckFunc(
					testAccInputCheck("graylog_input.gelf_udp", &name),
					resource.TestCheckResourceAttr(
						"graylog_input.gelf_udp", "title", "GELF UDP"),
					resource.TestCheckResourceAttr(
						"graylog_input.gelf_udp", "global", "true"),
					resource.TestCheckResourceAttr(
						"graylog_input.gelf_udp", "type",
						"org.graylog2.inputs.gelf.udp.GELFUDPInput"),
					resource.TestCheckResourceAttr(
						"graylog_input.gelf_udp", "gelf_udp.#", "1"),
					resource.TestCheckResourceAttr(
						"graylog_input.gelf_udp", "gelf_udp.2558345342.port", "22201"),
					resource.TestCheckResourceAttr(
						"graylog_input.gelf_udp", "gelf_udp.2558345342.bind_address", "0.0.0.0"),
				),
			},
		},
	})
}

func testAccInputCheck(rn string, name *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[rn]
		if !ok {
			return fmt.Errorf("resource not found: %s", rn)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("input id not set")
		}

		client := testAccProvider.Meta().(*graylog.Client)

		response := new(struct {
			Title        string
			Global       bool
			Name         string
			ContentPack  *string `json:"content_pack,omitempty"`
			Type         string
			Attributes   map[string]interface{}
			StaticFields map[string]string `json:"static_fields"`
		})

		url := fmt.Sprintf("/api/system/inputs/%s", rs.Primary.ID)
		err := client.Get(url, response)
		if err != nil {
			return err
		}

		*name = response.Name

		return testutils.AssertGoldenJSON("test-fixtures/input-gelf-udp.golden", response)
	}
}

func testAccInputCheckDestroy(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*graylog.Client)

		response := new(struct {
			Inputs []struct {
				Name string
			}
		})

		err := client.Get("/api/system/inputs", response)
		if err != nil {
			return err
		}

		for _, in := range response.Inputs {
			if in.Name == name {
				return fmt.Errorf("input still exists after destroy")
			}
		}

		return nil
	}
}

const testAccInputConfig_gelfUDP = `
resource "graylog_input" "gelf_udp" {
  title  = "GELF UDP"
  global = true

  gelf_udp {
    bind_address = "0.0.0.0"
    port         = 22201
  }
}
`
