package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rebuy-de/terraform-provider-graylog/pkg/graylog"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"server_url": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "http://localhost:9000",
				Description: "URL to the Graylog API",
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "admin",
				Description: "Username for the Graylog API",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "http://localhost:9000",
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
