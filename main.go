package main

import (
	"cc_gatherer_gin/DB"
	"cc_gatherer_gin/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
)

func test() {
	user, err := DB.Users.Query("18085268536")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Print(user.Name)
}

func main() {
	// 打开数据库
	err := DB.SqlOpen()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// test()
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r = CollectRoute(r)
	panic(r.Run(":9000"))
}
