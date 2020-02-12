package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"graylog": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {
	for _, name := range []string{"GRAYLOG_SERVER_URL", "GRAYLOG_USERNAME", "GRAYLOG_PASSWORD"} {
		if v := os.Getenv(name); v == "" {
			t.Fatal("GRAYLOG_SERVER_URL, GRAYLOG_USERNAME and GRAYLOG_PASSWORD must be set for acceptance tests")
		}
	}
}
