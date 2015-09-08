package main

import "fmt"
import "github.com/t-iwano/wakame_api_client/api"

func main() {
     // initialize Api Client.
     a := api.NewApiInitialize()

     fmt.Println("Call Image Endpoint Methods.")
     a.Image().Index()
     a.Image().Show()     
     a.Image().Create()     
     a.Image().Update()     
     a.Image().Destroy()     
}