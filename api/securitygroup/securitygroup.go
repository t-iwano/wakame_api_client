package securitygroup

import "github.com/t-iwano/wakame_api_client/http"

type Endpoint struct {
	http *http.Client
}

func (endpoint *Endpoint) Index() {
	endpoint.http.SendRequest("GET", "security_groups", nil)
}

func (endpoint *Endpoint) Show(id string) {
	endpoint.http.SendRequest("GET", "security_groups", id)
}

func (endpoint *Endpoint) Create(params map[string]string) {
	endpoint.http.SendRequest("POST", "security_groups", params)
}

func (endpoint *Endpoint) Update(params map[string]string) {
	endpoint.http.SendRequest("PUT", "security_groups", params)
}

func (endpoint *Endpoint) Destroy(id string) {
	endpoint.http.SendRequest("DELETE", "security_groups", id)
}

func NewSecuritygroupInitialize(client *http.Client) *Endpoint {
	return &Endpoint{
		http: client,
	}
}
