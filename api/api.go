package api

import "github.com/t-iwano/wakame_api_client/api/image"

type Api struct {
     image *image.Endpoint
}

func (api *Api) Image() *image.Endpoint {
     return api.image
}

func NewApiInitialize() *Api {
     return &Api{
     	    image.NewImageInitialize(),
     }
}
