package users

import (
	"github.com/gin-gonic/gin"

	"github.com/mayur-ralali/tachyon/sdk/lib"
	"github.com/mayur-tolexo/flash"
)

type UserClient struct {
	flash.Server `v:"1" root:"/test/"`
	add          flash.GET `url:"/add/:a/:b/"`
}

func (*UserClient) Add(c *gin.Context) {
	a, _ := lib.GetParamInt64(c, "a")
	b, _ := lib.GetParamInt64(c, "b")
	lib.BuildResponse(c, a+b, nil)
}
