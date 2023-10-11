package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Set the router as the default one shipped with Gin
	r := gin.Default()

	// 使用Cors中间件来配置CORS
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//Getenv CORS_SITE, and split it with comma, and assign it to allowOrigins
	corsSite := os.Getenv("CORS_SITE")
	//filter the corsSite, and remove the space
	corsSite = strings.ReplaceAll(corsSite, " ", "")
	config := cors.DefaultConfig()
	config.AllowOrigins = strings.Split(corsSite, ",")
	fmt.Println(config.AllowOrigins)
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}

	r.Use(cors.New(config))
	r.Use(static.Serve("/", static.LocalFile("./client/dist", true)))

	api := r.Group("/api")
	{

		api.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "This is a response from the server.",
			})
		})
	}
	r.Run(":5000")
}
