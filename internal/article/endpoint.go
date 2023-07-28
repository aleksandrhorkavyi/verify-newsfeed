package article

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	gin *gin.Engine
	svc Service
}

func NewController(g *gin.Engine, svc Service) *Controller {
	return &Controller{gin: g, svc: svc}
}

func (c *Controller) InitRoutes() {
	c.gin.GET("/articles", c.all)
	c.gin.GET("/articles/:id", c.one)

}

func (e *Controller) all(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func (e *Controller) one(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}
