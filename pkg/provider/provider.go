package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/rebuy-de/terraform-provider-graylog/pkg/graylog"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"server_url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("GRAYLOG_SERVER_URL", "http://localhost:9000"),
				Description: "URL to the Graylog API",
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("GRAYLOG_USERNAME", "admin"),
				Description: "Username for the Graylog API",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("GRAYLOG_PASSWORD", "admin"),
				Description: "Password for the Graylog API",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"graylog_input": resourceGraylogInput(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	return &graylog.Client{
		ServerURL: d.Get("server_url").(string),
		Username:  d.Get("username").(string),
		Password:  d.Get("password").(string),
	}, nil
}
