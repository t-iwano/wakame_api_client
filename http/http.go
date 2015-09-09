package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	defaultScheme    = "http"
	defaultHost      = "localhost"
	defaultPort      = "9001"
	defaultPath      = "/api/12.03"
	defaultAccountId = "a-shpoolxx"
)

func generateBaseUrl(endpoint string, id interface{}) string {
	baseurl := defaultScheme + "://" + defaultHost + ":" + defaultPort + defaultPath + "/" + endpoint

	switch id.(type) {
	case string:
		baseurl = baseurl + "/" + id.(string)
	}
	return baseurl
}

func generateNewRequest(verb string, baseurl string, data url.Values) (req *http.Request) {
	switch verb {
	case "GET", "DELETE":
		req, _ := http.NewRequest(verb, baseurl, nil)
		return req
	case "POST", "PUT":
		req, _ := http.NewRequest(verb, baseurl, strings.NewReader(data.Encode()))
		return req
	}
	return
}

func SendRequest(verb string, suffix string, params interface{}) {

	baseurl := generateBaseUrl(suffix, params)

	data := url.Values{}

	switch params.(type) {
	case map[string]string:
	     fmt.Println(params)
	}

	req := generateNewRequest(verb, baseurl, data)

	if verb == "GET" && data != nil {
		req.URL.RawQuery = data.Encode()
	}

	req.Header.Set("X_VDC_ACCOUNT_UUID", defaultAccountId)

	client := new(http.Client)
	res, _ := client.Do(req)
	defer res.Body.Close()

	contents, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(contents))

}

