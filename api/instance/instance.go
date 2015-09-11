package instance

import "github.com/t-iwano/wakame_api_client/http"

type Endpoint struct {
	http *http.Client
}

func (endpoint *Endpoint) Index() {
	endpoint.http.SendRequest("GET", "instances", nil)
}

func (endpoint *Endpoint) Show(id string) {
	endpoint.http.SendRequest("GET", "instances", id)
}

func (endpoint *Endpoint) Create(params map[string]string) {
	endpoint.http.SendRequest("POST", "instances", params)
}

func (endpoint *Endpoint) Update(params map[string]string) {
	endpoint.http.SendRequest("PUT", "instances", params)
}

func (endpoint *Endpoint) Destroy(id string) {
	endpoint.http.SendRequest("DELETE", "instances", id)
}

func NewInstanceInitialize(client *http.Client) *Endpoint {
	return &Endpoint{
		http: client,
	}
}
