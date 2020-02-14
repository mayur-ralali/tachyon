package main

import (
	"os"

	"github.com/mayur-ralali/tachyon/account/api/users"
	"github.com/mayur-tolexo/flash"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
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
