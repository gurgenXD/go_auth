package routers

import (
	v1 "auth/api/routers/v1"

	"github.com/gin-gonic/gin"
)

const (
	APIV1Path = "/api/v1"
)

func New() (router *gin.Engine) {
	router = gin.Default()

	v1Router := router.Group(APIV1Path)
	{
		v1Router.GET("", v1.Ping)
	}

	return
}
