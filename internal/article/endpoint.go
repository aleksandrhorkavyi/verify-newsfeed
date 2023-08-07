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
	c.gin.GET("/articles/:id", c.one)
	c.gin.GET("/do-something", c.all)
}

func (c *Controller) all(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, c.svc.DoSomething())
}

func (c *Controller) one(ctx *gin.Context) {
	id := ctx.Param("id")

	ctx.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}
