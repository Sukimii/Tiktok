package routers

import (
	"github.com/gin-gonic/gin"
	"tiktok/controller"
	"tiktok/middleware"
)

func BaseRoutersInit(r *gin.Engine) {
	baseRouters := r.Group("/douyin")
	{
		baseRouters.GET("/feed", middleware.JWTMiddleWare(), controller.Feed)
		baseRouters.POST("/publish/action/", middleware.JWTMiddleWare(), controller.PublishVideoController)
	}

}
