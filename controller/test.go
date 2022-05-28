package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Stu struct {
	Name   string `json:"name"`
	Number string `json:"number"`
}

func Test(ctx *gin.Context) {
	stuS := ctx.PostForm("stu")
	stu := &Stu{}
	err := json.Unmarshal([]byte(stuS), &stu)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusServiceUnavailable,
			"msg":  "解析错误",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"stu":  stu,
	})
}
