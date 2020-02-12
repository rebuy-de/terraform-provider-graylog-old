package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/rebuy-de/terraform-provider-graylog/pkg/graylog"
	"github.com/rebuy-de/terraform-provider-graylog/pkg/testutils"
)

func TestAccInput_gelfUDP(t *testing.T) {
	var title string
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccInputCheckDestroy(title),
		Steps: []resource.TestStep{
			{
				Config: testAccInputConfig_gelfUDP,
				Check: resource.ComposeTestCheckFunc(
					testAccInputCheck("graylog_input.gelf_udp", &title, true),
					resource.TestCheckResourceAttr(
						"graylog_input.gelf_udp", "title", "gelf-udp"),
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

func TestAccInput_gelfTCP(t *testing.T) {
	var title string
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccInputCheckDestroy(title),
		Steps: []resource.TestStep{
			{
				Config: testAccInputConfig_gelfTCP,
				Check: resource.ComposeTestCheckFunc(
					testAccInputCheck("graylog_input.gelf_tcp", &title, true),
					resource.TestCheckResourceAttr(
						"graylog_input.gelf_tcp", "title", "gelf-tcp"),
					resource.TestCheckResourceAttr(
						"graylog_input.gelf_tcp", "global", "true"),
					resource.TestCheckResourceAttr(
						"graylog_input.gelf_tcp", "type",
						"org.graylog2.inputs.gelf.tcp.GELFTCPInput"),
					resource.TestCheckResourceAttr(
						"graylog_input.gelf_tcp", "gelf_tcp.#", "1"),
					resource.TestCheckResourceAttr(
						"graylog_input.gelf_tcp", "gelf_tcp.518599376.port", "12201"),
					resource.TestCheckResourceAttr(
						"graylog_input.gelf_tcp", "gelf_tcp.518599376.bind_address", "0.0.0.0"),
				),
			},
		},
	})
}

func TestAccInput_gelfHTTP(t *testing.T) {
	var title string
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccInputCheckDestroy(title),
		Steps: []resource.TestStep{
			{
				Config: testAccInputConfig_gelfHTTP,
				Check: resource.ComposeTestCheckFunc(
					testAccInputCheck("graylog_input.gelf_http", &title, true),
					resource.TestCheckResourceAttr(
						"graylog_input.gelf_http", "title", "gelf-http"),
					resource.TestCheckResourceAttr(
						"graylog_input.gelf_http", "global", "true"),
					resource.TestCheckResourceAttr(
						"graylog_input.gelf_http", "type",
						"org.graylog2.inputs.gelf.http.GELFTCPInput"),
					resource.TestCheckResourceAttr(
						"graylog_input.gelf_http", "gelf_http.#", "1"),
					resource.TestCheckResourceAttr(
						"graylog_input.gelf_http", "gelf_http.518599376.port", "12202"),
					resource.TestCheckResourceAttr(
						"graylog_input.gelf_http", "gelf_tcp.518599376.bind_address", "0.0.0.0"),
				),
			},
		},
	})
}

func TestAccInput_beats(t *testing.T) {
	var title string
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccInputCheckDestroy(title),
		Steps: []resource.TestStep{
			{
				Config: testAccInputConfig_beats,
				Check: resource.ComposeTestCheckFunc(
					testAccInputCheck("graylog_input.beats", &title, true),
					resource.TestCheckResourceAttr(
						"graylog_input.beats", "title", "beats"),
					resource.TestCheckResourceAttr(
						"graylog_input.beats", "global", "true"),
					resource.TestCheckResourceAttr(
						"graylog_input.beats", "type",
						"org.graylog.plugins.beats.BeatsInput"),
					resource.TestCheckResourceAttr(
						"graylog_input.beats", "beats.#", "1"),
					resource.TestCheckResourceAttr(
						"graylog_input.beats", "beats.941094508.port", "5044"),
					resource.TestCheckResourceAttr(
						"graylog_input.beats", "beats.941094508.bind_address", "0.0.0.0"),
				),
			},
		},
	})
}

func TestAccInput_changeAddress(t *testing.T) {
	var title string
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccInputCheckDestroy(title),
		Steps: []resource.TestStep{
			{
				Config: testAccInputConfig_gelfTCP,
				Check: resource.ComposeTestCheckFunc(
					testAccInputCheck("graylog_input.gelf_tcp", &title, false),
					resource.TestCheckResourceAttr(
						"graylog_input.gelf_tcp", "gelf_tcp.518599376.bind_address", "0.0.0.0"),
					resource.TestCheckResourceAttr(
						"graylog_input.gelf_tcp", "gelf_tcp.518599376.port", "12201"),
				),
			},
			{
				Config: testAccInputConfig_gelfTCP_changedAddress,
				Check: resource.ComposeTestCheckFunc(
					testAccInputCheck("graylog_input.gelf_tcp", &title, false),
					resource.TestCheckResourceAttr(
						"graylog_input.gelf_tcp", "gelf_tcp.1984762550.bind_address", "127.0.0.1"),
					resource.TestCheckResourceAttr(
						"graylog_input.gelf_tcp", "gelf_tcp.1984762550.port", "12201"),
					resource.TestCheckNoResourceAttr(
						"graylog_input.gelf_tcp", "gelf_tcp.518599376.bind_address"),
					resource.TestCheckNoResourceAttr(
						"graylog_input.gelf_tcp", "gelf_tcp.518599376.port"),
				),
			},
		},
	})
}

func testAccInputCheck(rn string, title *string, golden bool) resource.TestCheckFunc {
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

		*title = response.Title

		if golden {
			return testutils.AssertGoldenJSON(fmt.Sprintf("test-fixtures/input-%s.golden", *title), response)
		}

		return nil
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

func TestAccInput_multiBlock(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccInputConfig_multiBlock,
				ExpectError: regexp.MustCompile("graylog_input expects exactly one block of "),
			},
		},
	})
}

const testAccInputConfig_gelfUDP = `
resource "graylog_input" "gelf_udp" {
  title  = "gelf-udp"
  global = true

  gelf_udp {
    bind_address = "0.0.0.0"
    port         = 22201
  }
}
`
const testAccInputConfig_gelfTCP = `
resource "graylog_input" "gelf_tcp" {
  title  = "gelf-tcp"
  global = true

  gelf_tcp {
    bind_address = "0.0.0.0"
    port         = 12201
  }
}
`

const testAccInputConfig_gelfHTTP = `
resource "graylog_input" "gelf_http" {
  title  = "gelf-http"
  global = true

  gelf_http {
    bind_address = "0.0.0.0"
    port         = 12201
  }
}
`

const testAccInputConfig_gelfTCP_changedAddress = `
resource "graylog_input" "gelf_tcp" {
  title  = "gelf-tcp"
  global = true

  gelf_tcp {
    bind_address = "127.0.0.1"
    port         = 12201
  }
}
`

const testAccInputConfig_beats = `
resource "graylog_input" "beats" {
  title  = "beats"
  global = true

  beats {
    bind_address = "0.0.0.0"
    port         = 5044
  }
}
`

const testAccInputConfig_multiBlock = `
resource "graylog_input" "gelf_both" {
  title  = "gelf-both"
  global = true

  gelf_udp {
    bind_address = "0.0.0.0"
    port         = 22201
  }

  gelf_tcp {
    bind_address = "0.0.0.0"
    port         = 12201
  }
}

`
