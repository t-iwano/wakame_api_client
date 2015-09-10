package instance

import "github.com/t-iwano/wakame_api_client/http"

type Endpoint struct {
}

func (endpoint *Endpoint) Index() {
     http.SendRequest("GET", "instances", nil)
}

func (endpoint *Endpoint) Show(id string) {
     http.SendRequest("GET", "instances", id)
}

func (endpoint *Endpoint) Create(params map[string]string) {
     http.SendRequest("POST", "instances", params)
}

func (endpoint *Endpoint) Update(params map[string]string) {
     http.SendRequest("PUT", "instances", params)
}

func (endpoint *Endpoint) Destroy(id string) {
     http.SendRequest("DELETE", "instances", id)
}

func NewInstanceInitialize() *Endpoint {
     return &Endpoint{}
}
