package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//GinRouter -> Gin Router
type GinRouter struct {
	Gin *gin.Engine
}

//NewGinRouter all the routes are defined here
func NewGinRouter() GinRouter {

	httpRouter := gin.Default()

	httpRouter.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Polls Up and Running..."})
	})
	return GinRouter{
		Gin: httpRouter,
	}

}
