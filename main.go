// basic server
package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/krishpranav/golang-management/middleware"
	"github.com/krishpranav/golang-management/routers"
)

// server is running[now we can able to create users]
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routers.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.Run(":" + port)
}
