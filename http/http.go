package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/t-iwano/wakame_api_client/apiconfig"
)

const (
	defaultScheme    = "http"
	defaultHost      = "localhost"
	defaultPort      = "9001"
	defaultPath      = "/api/12.03"
	defaultAccountId = "a-shpoolxx"
)

type Client struct {
	config *apiconfig.WakameConfig
}

func (c *Client) generateBaseUrl(endpoint string, params interface{}) string {
	var scheme string
	if c.config.Scheme == "" {
		scheme = defaultScheme
	} else {
		scheme = c.config.Scheme
	}

	var host string
	if c.config.Host == "" {
		host = defaultHost
	} else {
		host = c.config.Host
	}

	var port string
	if c.config.Port == "" {
		port = defaultPort
	} else {
		port = c.config.Port
	}

	var path string
	if c.config.Path == "" {
		path = defaultPath
	} else {
		path = c.config.Path
	}

	baseurl := scheme + "://" + host + ":" + port + path + "/" + endpoint

	switch t := params.(type) {
	case string:
		baseurl = baseurl + "/" + t
	case map[string]string:
		if id, ok := t["id"]; ok {
			baseurl = baseurl + "/" + id
		}
	}
	return baseurl
}

func (c *Client) generateNewRequest(verb string, baseurl string, data url.Values) (req *http.Request) {
	switch verb {
	case "GET", "DELETE", "PUT":
		req, _ := http.NewRequest(verb, baseurl, nil)
		return req
	case "POST":
		req, _ := http.NewRequest(verb, baseurl, strings.NewReader(data.Encode()))
		return req
	}
	return
}

func (c *Client) SendRequest(verb string, suffix string, params interface{}) {

	baseurl := c.generateBaseUrl(suffix, params)

	data := url.Values{}

	switch p := params.(type) {
	case map[string]string:
		for k, v := range p {
			data.Add(k, v)
		}
	}

	req := c.generateNewRequest(verb, baseurl, data)

	if (verb == "GET" || verb == "PUT") && data != nil {
		req.URL.RawQuery = data.Encode()
	}

	var accountid string
	if c.config.AccountId == "" {
		accountid = defaultAccountId
	} else {
		accountid = c.config.AccountId
	}
	req.Header.Set("X_VDC_ACCOUNT_UUID", accountid)

	client := new(http.Client)
	res, _ := client.Do(req)
	defer res.Body.Close()

	contents, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(contents))

}

func (c *Client) Config() *apiconfig.WakameConfig {
	return c.config
}

func NewHttpInitialize(config *apiconfig.WakameConfig) *Client {
	return &Client{
		config: config,
	}
}
