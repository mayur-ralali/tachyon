package main

import (
	"github.com/mayur-ralali/tachyon/account/endpoint"
	"github.com/mayur-tolexo/flash"
)

func main() {
	router := flash.Default()
	router.AddService(&endpoint.UserClient{})
	router.Start(":8005")
}
