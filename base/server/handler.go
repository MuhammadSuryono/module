package server

import (
	"fmt"
	"github.com/MuhammadSuryono/module/base/database"
	"github.com/MuhammadSuryono/module/base/response"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"time"
)

func CreateHttpServer() *gin.Engine {

	errorEnv := godotenv.Load()
	if errorEnv != nil {
		fmt.Println("Error loading .env file")
	}

	database.CreateConnection()

	server := gin.Default()

	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
		MaxAge: 12 * time.Hour,
	}))

	server.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, response.CommonResponse{
			IsSuccess: false,
			Message: "Route not found",
		})
	})

	server.GET("/", serviceInfo(
		os.Getenv("APP_NAME"),
		os.Getenv("VERSION"),
		"TEAM_BACKEND_OKTAPOS"))

	return server
}

func serviceInfo(app string, version string, author string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, response.CommonResponse{
			IsSuccess: false,
			Message: "Service Info",
			Data: map[string]interface{}{
				"app_name": app,
				"version": version,
				"author": author,
			},
		})
	}
}