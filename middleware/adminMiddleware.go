package middleware

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mayur-ralali/tachyon/sdk/objects"
)

func AdminMiddleWare() gin.HandlerFunc {

	return func(c *gin.Context) {
		responseHelper := helpers.ResponseHelper{}
		oldAuthorization := c.GetHeader("x-access-token")

		if oldAuthorization == "" {
			response := responseHelper.UnauthorizedResponse()
			c.JSON(response.Error, response)
			c.Abort()
			return
		}

		url := fmt.Sprintf("%s/v1/profile", os.Getenv("OAUTH_SERVER_URL"))

		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Add("x-access-token", oldAuthorization)
		req.Header.Add("cache-control", "no-cache")
		res, _ := http.DefaultClient.Do(req)

		if res.StatusCode == 200 {
			defer res.Body.Close()
			body, _ := ioutil.ReadAll(res.Body)

			var result objects.OauthUserProfileObject
			json.Unmarshal([]byte(string(body)), &result)
			if result.Success.UserSession.User.Type == "A" {
				c.Set("user", result)
				return
			}
		}

		response := responseHelper.ForbiddenResponse()
		c.JSON(response.Error, response)
		c.Abort()
	}
}
