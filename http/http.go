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

func generateBaseUrl(endpoint string, params interface{}) string {
	baseurl := defaultScheme + "://" + defaultHost + ":" + defaultPort + defaultPath + "/" + endpoint

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

func generateNewRequest(verb string, baseurl string, data url.Values) (req *http.Request) {
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

func SendRequest(verb string, suffix string, params interface{}) {

	baseurl := generateBaseUrl(suffix, params)

	data := url.Values{}

        switch p := params.(type) {
        case map[string]string:
                for k, v := range p {
                        data.Add(k, v)
                }
        }

	req := generateNewRequest(verb, baseurl, data)

	if (verb == "GET" || verb == "PUT") && data != nil {
		req.URL.RawQuery = data.Encode()
	}

	req.Header.Set("X_VDC_ACCOUNT_UUID", defaultAccountId)

	client := new(http.Client)
	res, _ := client.Do(req)
	defer res.Body.Close()

	contents, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(contents))

}

