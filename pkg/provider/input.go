package provider

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rebuy-de/terraform-provider-graylog/pkg/graylog"
	"github.com/rebuy-de/terraform-provider-graylog/pkg/types"
)

func resourceGraylogInput() *schema.Resource {
	return &schema.Resource{
		Create: resourceGraylogInputCreate,
		Read:   resourceGraylogInputRead,
		Update: resourceGraylogInputUpdate,
		Delete: resourceGraylogInputDelete,

		Schema: map[string]*schema.Schema{
			"title": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"global": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"node": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gelf_udp": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,

				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},
						"bind_address": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Default:  "0.0.0.0",
						},
					},
				},
			},
		},
	}
}

func resourceGraylogInputCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*graylog.Client)

	request, err := resourceGraylogInputGenerateCreateRequest(d)
	if err != nil {
		return err
	}

	response := new(types.SystemInputCreateResponse)
	err = client.Post("/api/system/inputs", request, response)
	if err != nil {
		return err
	}

	d.SetId(response.ID)

	return nil
}

func resourceGraylogInputRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*graylog.Client)

	input := new(types.SystemInputSummary)
	url := fmt.Sprintf("/api/system/inputs/%s", d.Id())
	err := client.Get(url, input)
	if err != nil {
		return err
	}

	d.Set("title", input.Title)

	return nil
}

func resourceGraylogInputUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*graylog.Client)

	request, err := resourceGraylogInputGenerateCreateRequest(d)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("/api/system/inputs/%s", d.Id())
	return client.Put(url, request, nil)
}

func resourceGraylogInputDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*graylog.Client)
	url := fmt.Sprintf("/api/system/inputs/%s", d.Id())
	return client.Delete(url)
}

func resourceGraylogInputGenerateCreateRequest(d *schema.ResourceData) (*types.SystemInputCreateRequest, error) {
	request := &types.SystemInputCreateRequest{
		Title:  d.Get("title").(string),
		Global: d.Get("global").(bool),
		Node:   d.Get("node").(string),
	}

	gelfUDP := d.Get("gelf_udp").(*schema.Set)

	blockCount := 0
	blockCount += gelfUDP.Len()

	if blockCount != 1 {
		return nil, fmt.Errorf("graylog_input expects exactly one 'gelf_udp' block")
	}

	switch {
	case gelfUDP.Len() == 1:
		request.Type = "org.graylog2.inputs.gelf.udp.GELFUDPInput"
		request.Configuration = gelfUDP.List()[0].(map[string]interface{})
	default:
		// shouldn't happen, because the blockCount got checked before
		panic("unexpected block count")
	}

	return request, nil
}
