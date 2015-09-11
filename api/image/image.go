package image

import "fmt"
import "github.com/t-iwano/wakame_api_client/http"

type Endpoint struct {
	http *http.Client
}

func (endpoint *Endpoint) Index() {
	endpoint.http.SendRequest("GET", "images", nil)
}

func (endpoint *Endpoint) Show(id string) {
	endpoint.http.SendRequest("GET", "images", id)
}

func (endpoint *Endpoint) Create() {
	fmt.Println("Create new Image.")
}

func (endpoint *Endpoint) Update() {
	fmt.Println("Update Image.")
}

func (endpoint *Endpoint) Destroy() {
	fmt.Println("Delete Image.")
}

func NewImageInitialize(client *http.Client) *Endpoint {
	return &Endpoint{
		http: client,
	}
}
