package main

import (
	"cc_gatherer_gin/controller"
	"cc_gatherer_gin/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})
	r.POST("/user/register", controller.Register)
	r.POST("/user/login", controller.Login)
	r.GET("/user/info", middleware.AuthMiddleware(), controller.Info)
	r.POST("/user/words", controller.Words)
	r.POST("/test", controller.Test)
	r.POST("/fanyi", controller.Fanyi)
	return r
}
