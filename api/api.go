package api

import "github.com/t-iwano/wakame_api_client/api/image"
import "github.com/t-iwano/wakame_api_client/api/securitygroup"

type Api struct {
     image *image.Endpoint
     securitygroup *securitygroup.Endpoint
}

func (api *Api) Image() *image.Endpoint {
     return api.image
}

func (api *Api) Securitygroup() *securitygroup.Endpoint {
     return api.securitygroup
}

func NewApiInitialize() *Api {
     return &Api{
     	    image.NewImageInitialize(),
     	    securitygroup.NewSecuritygroupInitialize(),
     }
}
