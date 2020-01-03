package app

import (
	"github.com/gin-gonic/gin"
	"github.com/janhaans/articles/controller"
)

func mapURL(router *gin.Engine) {
	router.GET("/", controller.ShowIndexPage)
}
