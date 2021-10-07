package main

import (
	"cc_gatherer_gin/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r = CollectRoute(r)
	panic(r.Run(":9000"))
}
