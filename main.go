package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	_ "swagger_setup/docs"
)

// @title Simple Gin HTTP Request API
// @version 1.0
// @description This is a simple API to demonstrate Swagger with Gin in Go.

// @host localhost:8080
// @BasePath /api/v1

// SimpleRequest godoc
// @Summary Show a simple request example
// @Description get a simple HTTP response
// @Tags example
// @Accept  json
// @Produce  json
// @Success 200 {string} string "ok"
// @Router /simple [get]
func SimpleRequest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, this is a simple request!",
	})
}

func main() {
	// Initialize Gin router
	r := gin.Default()

	// Swagger route

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API routes
	v1 := r.Group("/api/v1")
	{
		v1.GET("/simple", SimpleRequest)
	}

	// Start the server
	r.Run(":8080")
}
