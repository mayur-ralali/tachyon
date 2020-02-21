package main

import (
	"os"

	"github.com/mayur-ralali/tachyon/account/api/users"
	"github.com/mayur-ralali/tachyon/sdk/lib"
	"github.com/mayur-tolexo/flash"
)

func init() {
	lib.LoadEnv()
}

func main() {
	router := flash.Default()
	addServices(router)
	host := os.Getenv("ACCOUNT_HOST")
	port := os.Getenv("ACCOUNT_PORT")
	router.Start(host + ":" + port)
}

func addServices(router *flash.Server) {
	router.AddService(&users.UserClient{})
}
