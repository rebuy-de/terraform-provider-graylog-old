package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/pkg/errors"

	"github.com/rebuy-de/terraform-provider-graylog/pkg/graylog"
	"github.com/rebuy-de/terraform-provider-graylog/pkg/types"
)

func resourceGraylogInput() *schema.Resource {
	return &schema.Resource{
		Create: resourceGraylogInputCreate,
		Read:   resourceGraylogInputRead,
		Update: resourceGraylogInputUpdate,
		Delete: resourceGraylogInputDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"title": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"global": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"node": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"gelf_udp": {
				Type:     schema.TypeSet,
				Optional: true,

				Elem: &schema.Resource{
					// I didn't find any documentation about the GELF UDP input configuration.
					// Therefore we have to add options as needed.
					Schema: map[string]*schema.Schema{
						"port": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"bind_address": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "0.0.0.0",
						},
					},
				},
			},
			"gelf_tcp": {
				Type:     schema.TypeSet,
				Optional: true,

				Elem: &schema.Resource{
					// I didn't find any documentation about the GELF TCP input configuration.
					// Therefore we have to add options as needed.
					Schema: map[string]*schema.Schema{
						"port": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"bind_address": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "0.0.0.0",
						},
					},
				},
			},
			"beats": {
				Type:     schema.TypeSet,
				Optional: true,

				Elem: &schema.Resource{
					// I didn't find any documentation about the beats input configuration.
					// Therefore we have to add options as needed.
					Schema: map[string]*schema.Schema{
						"port": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"bind_address": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "0.0.0.0",
						},
					},
				},
			},
			"gelf_http": {
				Type:     schema.TypeSet,
				Optional: true,

				Elem: &schema.Resource{
					// I didn't find any documentation about the GELF Http input configuration.
					// Therefore we have to add options as needed.
					Schema: map[string]*schema.Schema{
						"port": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"bind_address": {
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

	return resourceGraylogInputRead(d, meta)
}

func resourceGraylogInputRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*graylog.Client)

	input := new(types.SystemInputSummary)
	url := fmt.Sprintf("/api/system/inputs/%s", d.Id())
	err := client.Get(url, input)
	_ = err // TODO: Gracefully handle 404s

	d.Set("title", input.Title)
	d.Set("global", input.Global)
	d.Set("name", input.Name)
	d.Set("type", input.Type)
	d.Set("attributes", input.Attributes)
	d.Set("static_fields", input.StaticFields)

	switch input.Type {
	case "org.graylog2.inputs.gelf.udp.GELFUDPInput":
		d.Set("gelf_udp", []interface{}{input.Attributes})
	case "org.graylog2.inputs.gelf.tcp.GELFTCPInput":
		d.Set("gelf_tcp", []interface{}{input.Attributes})
	case "org.graylog.plugins.beats.BeatsInput":
		d.Set("beats", []interface{}{input.Attributes})
	case "org.graylog2.inputs.gelf.http.GELFHttpInput":
		d.Set("gelf_http", []interface{}{
			map[string]interface{}{
				"port":         input.Attributes["port"],
				"bind_address": input.Attributes["bind_address"],
			},
		})
	default:
		return errors.Errorf("unknown type %s", input.Type)
	}

	return nil
}

func resourceGraylogInputUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*graylog.Client)

	request, err := resourceGraylogInputGenerateCreateRequest(d)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("/api/system/inputs/%s", d.Id())
	err = client.Put(url, request, nil)
	if err != nil {
		return err
	}

	return resourceGraylogInputRead(d, meta)
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

	var (
		beats    = d.Get("beats").(*schema.Set)
		gelfHTTP = d.Get("gelf_http").(*schema.Set)
		gelfTCP  = d.Get("gelf_tcp").(*schema.Set)
		gelfUDP  = d.Get("gelf_udp").(*schema.Set)
	)

	blockCount := 0
	blockCount += beats.Len()
	blockCount += gelfHTTP.Len()
	blockCount += gelfTCP.Len()
	blockCount += gelfUDP.Len()

	if blockCount != 1 {
		return nil, fmt.Errorf("graylog_input expects exactly one block of 'gelf_udp', 'gelf_tcp' or 'beats'")
	}

	switch {
	case gelfUDP.Len() == 1:
		request.Type = "org.graylog2.inputs.gelf.udp.GELFUDPInput"
		request.Configuration = gelfUDP.List()[0].(map[string]interface{})
	case gelfTCP.Len() == 1:
		request.Type = "org.graylog2.inputs.gelf.tcp.GELFTCPInput"
		request.Configuration = gelfTCP.List()[0].(map[string]interface{})
	case beats.Len() == 1:
		request.Type = "org.graylog.plugins.beats.BeatsInput"
		request.Configuration = beats.List()[0].(map[string]interface{})
	case gelfHTTP.Len() == 1:
		request.Type = "org.graylog2.inputs.gelf.http.GELFHttpInput"
		request.Configuration = gelfHTTP.List()[0].(map[string]interface{})
	default:
		// shouldn't happen, because the blockCount got checked before
		panic("unexpected block count")
	}

	return request, nil
}
