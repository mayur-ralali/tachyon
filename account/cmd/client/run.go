package main

import (
	"github.com/mayur-ralali/tachyon/account/endpoint/users"
	"github.com/mayur-ralali/tachyon/account/constant"
	"github.com/mayur-tolexo/flash"
)

func main() {
	router := flash.Default()
	router.AddService(&users.UserClient{})
	router.Start(constant.ClientHost)
}
