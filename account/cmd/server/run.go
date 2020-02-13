package main

import (
	"log"
	"net"

	"github.com/mayur-ralali/tachyon/account/pb"
	"github.com/mayur-ralali/tachyon/account/server/users"
	"google.golang.org/grpc"
)

func main() {
	srv := grpc.NewServer()
	pb.RegisterAddServiceServer(srv, &users.Service{})
	l, err := net.Listen("tcp", ":8083")
	if err != nil {
		log.Fatalln(err)
	}
	srv.Serve(l)
}
