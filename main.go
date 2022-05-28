package main

import (
	"cc_gatherer_gin/DB"
	"cc_gatherer_gin/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// 打开数据库
	err := DB.SqlOpen()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r = CollectRoute(r)
	panic(r.Run(":9000"))
}
