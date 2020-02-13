package lib

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
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
	c.JSON(200, gin.H{
		"data":  data,
		"error": err,
	})
}
