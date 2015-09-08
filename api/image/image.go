package image

import "fmt"

type Endpoint struct {
}

func (endpoint *Endpoint) Index() {
     fmt.Println("Show image list.")
}

func (endpoint *Endpoint) Show() {
     fmt.Println("Show image.")
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

func NewImageInitialize() *Endpoint {
     return &Endpoint{}
}