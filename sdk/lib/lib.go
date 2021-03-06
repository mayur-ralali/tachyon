package lib

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
	"google.golang.org/grpc"
)

func ConnGrpc(opts ...grpc.DialOption) (conn *grpc.ClientConn, err error) {
	conn, err = grpc.Dial(":8083", opts...)
	return
}

func Fatalf(msg string, err error) {
	if err != nil {
		log.Fatalf(msg, err)
	}
}

//GetParamInt64 will return int value from param
func GetParamInt64(c *gin.Context, key string) (val int64, err error) {
	data := c.Param(key)
	val, err = strconv.ParseInt(data, 10, 64)
	return
}

//BuildResponse will build response
func BuildResponse(c *gin.Context, data interface{}, err error) {
	resp := gin.H{}
	if err != nil {
		resp["error"] = err
	} else {
		resp["data"] = data
	}
	c.JSON(200, resp)
}

//LoadEnv will load env
func LoadEnv() {
	priority := []string{
		".env.production",
		".env.staging",
		".env.dev",
		".env",
	}

	flag := false
	for _, curEnv := range priority {
		if err := gotenv.Load(curEnv); err == nil {
			flag = true
			break
		}
	}
	if !flag {
		log.Fatal("No env found")
	}
}
