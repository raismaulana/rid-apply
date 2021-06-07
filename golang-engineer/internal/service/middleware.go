package service

import (
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/raismaulana/rid-apply/golang-engineer/internal/config"
)

func AuthzGoogleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		access_token := authHeader[len(BEARER_SCHEMA):]
		resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + url.QueryEscape(access_token))
		if err != nil {
			config.Log.Error("Get: " + err.Error() + "\n")
			c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}
		defer resp.Body.Close()

		response, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			config.Log.Error("ReadAll: " + err.Error() + "\n")
			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}

		config.Log.Info("parseResponseBody: " + string(response) + "\n")
		c.Next()
	}
}
