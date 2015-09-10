package securitygroup

import "github.com/t-iwano/wakame_api_client/http"

type Endpoint struct {
}

func (endpoint *Endpoint) Index() {
     http.SendRequest("GET", "security_groups", nil)
}

func (endpoint *Endpoint) Show(id string) {
     http.SendRequest("GET", "security_groups", id)
}

func (endpoint *Endpoint) Create(params map[string]string) {
     http.SendRequest("POST", "security_groups", params)
}

func (endpoint *Endpoint) Update(params map[string]string) {
     http.SendRequest("PUT", "security_groups", params)
}

func (endpoint *Endpoint) Destroy(id string) {
     http.SendRequest("DELETE", "security_groups", id)
}

func NewSecuritygroupInitialize() *Endpoint {
     return &Endpoint{}
}