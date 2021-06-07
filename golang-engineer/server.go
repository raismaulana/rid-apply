package main

import (
	"github.com/gin-gonic/gin"
	"github.com/raismaulana/rid-apply/golang-engineer/internal/config"
	"github.com/raismaulana/rid-apply/golang-engineer/internal/service"
)

func main() {
	// Initialize Viper across the application
	config.InitializeViper()

	// Initialize Logger across the application
	config.InitializeZapCustomLogger()

	// db
	config.SetupDatabaseConnection()
	defer config.CloseDatabaseConnection(config.Db)

	// Initialize Oauth2 Services
	service.InitializeOAuthGoogle()

	// Routes for the application
	r := gin.New()
	r.GET("/", service.HandleMain)
	r.GET("/login-gl", service.HandleGoogleLogin)
	r.GET("/callback-gl", service.CallBackFromGoogle)
	r.Run()
}
