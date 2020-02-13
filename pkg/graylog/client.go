package graylog

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

type Client struct {
	Username  string
	Password  string
	ServerURL string
}

func (c *Client) do(method string, url string, reqValue, respValue interface{}) error {
	url = fmt.Sprintf("%s/%s",
		strings.TrimRight(c.ServerURL, "/"),
		strings.TrimLeft(url, "/"),
	)

	log.Debugf("%s %s", method, url)

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return err
	}

	req.SetBasicAuth(c.Username, c.Password)

	if reqValue != nil {
		buf := new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(reqValue)
		if err != nil {
			return err
		}

		log.Debug(buf.String())
		req.Body = ioutil.NopCloser(buf)
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode/100 != 2 {
		log.Debug(ReaderToString(resp.Body))
		return fmt.Errorf("Unexpected status %d", resp.StatusCode)
	}

	if respValue != nil {
		return json.NewDecoder(resp.Body).Decode(respValue)
	}

	return nil
}

func (c *Client) Get(url string, v interface{}) error {
	return c.do(GET, url, nil, v)
}

func (c *Client) Post(url string, reqValue, respValue interface{}) error {
	return c.do(POST, url, reqValue, respValue)
}

func (c *Client) Put(url string, reqValue, respValue interface{}) error {
	return c.do(PUT, url, reqValue, respValue)
}

func (c *Client) Delete(url string) error {
	return c.do(DELETE, url, nil, nil)
}
