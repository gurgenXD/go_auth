package routers

import (
	"github.com/gin-gonic/gin"
)

const (
	APIV1Path = "/api/v1"
)

func New() (router *gin.Engine) {
	router = gin.New()

	router.Use(gin.Recovery())

	v1Router := router.Group(APIV1Path)
	{
		v1Router.GET("", func(ctx *gin.Context) { ctx.JSON(200, nil) })
	}

	return
}
