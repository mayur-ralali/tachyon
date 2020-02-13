package endpoint

import (
	"github.com/gin-gonic/gin"
	"github.com/mayur-ralali/tachyon/account/pb"
	"github.com/mayur-ralali/tachyon/lib"
	"github.com/mayur-tolexo/flash"
	"google.golang.org/grpc"
)

type UserClient struct {
	flash.Server `v:"1" root:"/test/"`
	add          flash.GET `url:"/add/:a/:b/"`
}

func (*UserClient) Add(c *gin.Context) {

	conn, err := lib.ConnGrpc(grpc.WithInsecure())
	lib.Fatalf("could not create client: %v", err)

	a, _ := lib.GetParamInt64(c, "a")
	b, _ := lib.GetParamInt64(c, "b")

	client := pb.NewAddServiceClient(conn)
	req := &pb.Request{A: a, B: b}
	resp, err := client.Add(c, req)

	lib.BuildResponse(c, resp.Result, err)
}
