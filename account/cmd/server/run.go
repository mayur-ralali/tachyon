package main

import (
	"log"
	"net"

	"github.com/mayur-ralali/tachyon/account/pb"
	"github.com/mayur-ralali/tachyon/account/server/users"
	"google.golang.org/grpc"
	"github.com/mayur-ralali/tachyon/account/constant"
)

func main() {
	srv := grpc.NewServer()
	pb.RegisterAddServiceServer(srv, &users.Service{})
	l, err := net.Listen("tcp", constant.ServerHost)
	if err != nil {
		log.Fatalln(err)
	}
	srv.Serve(l)
}
