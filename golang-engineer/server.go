package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	google "github.com/zalando/gin-oauth2/google"
)

func main() {
	r := gin.Default()
	gSes := oauth2GoogleSetup()
	r.Use(google.Session(gSes))
	r.GET("/login", google.LoginHandler)
	// protected url group
	private := r.Group("/auth")
	private.Use(google.Auth())
	private.GET("/", UserInfoHandler)
	private.GET("/api", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Hello from private for groups"})
	})

	r.Run()
}

func oauth2GoogleSetup() string {
	redirectURL := "http://127.0.0.1:8080/auth/"
	credFile := "client_secret_854514465626-3k5lqp30e3cscs1tib655o5f54jnar5r.apps.googleusercontent.com.json" // you have to build your own

	scopes := []string{
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/userinfo.profile",
		"openid",
	}
	secret := []byte("secret") //
	sessionName := "goquestsession"

	// init settings for google auth
	google.Setup(redirectURL, credFile, scopes, secret)
	return sessionName
}
func UserInfoHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"Hello": "from private", "user": ctx.MustGet("user").(google.User)})
}
