package main

import (
	"blog/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"data": "Hello World !"})
	})
	router.Run(":8000")
	infrastructure.LoadEnv()
	infrastructure.NewDatabase()
}
