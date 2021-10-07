package controller

import (
	"cc_gatherer_gin/common"
	"cc_gatherer_gin/util"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx *gin.Context) {
	// 获取参数
	userName := ctx.PostForm("userName")
	userTel := ctx.PostForm("userTel")
	userEmail := ctx.PostForm("userEmail")
	userPassword := ctx.PostForm("userPassword")
	// 读取数据库
	users, err := util.ReadJsons("data/user.json")
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": http.StatusServiceUnavailable,
			"msg":  "用户数据读取出错",
		})
		return
	}
	// 检查手机号唯一性
	isExist := false
	for _, item := range users {
		if item["tel"] == userTel {
			isExist = true
		}
	}
	if isExist {
		ctx.JSON(200, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "手机号已存在",
		})
		return
	}
	pass, err := bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": http.StatusServiceUnavailable,
			"msg":  "系统错误，注册失败",
		})
	}

	// 添加到数据库
	users = append(users, gin.H{
		"name":     userName,
		"tel":      userTel,
		"email":    userEmail,
		"password": string(pass),
	})
	err = util.SaveJsons(users, "data/user.json")
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusServiceUnavailable,
			"msg":  "系统错误，注册失败",
		})
		return
	}
	// 获取Token
	token, err := common.ReleaseToken(userTel)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusServiceUnavailable,
			"msg":  "注册成功，但登录失败",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":  200,
		"token": token,
		"msg":   "注册成功",
	})
}

func Login(ctx *gin.Context) {
	// 获取参数
	userTel := ctx.PostForm("userTel")
	userPassword := ctx.PostForm("userPassword")
	// 读取数据库
	users, err := util.ReadJsons("data/user.json")
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": http.StatusServiceUnavailable,
			"msg":  "用户数据读取出错",
		})
		return
	}
	// 查找用户
	isExist := false
	user := gin.H{}
	for _, item := range users {
		if item["tel"] == userTel {
			isExist = true
			user = item
		}
	}
	// 如果用户不存在
	if !isExist {
		ctx.JSON(200, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "账号或密码错误，登录失败",
		})
		return
	}
	// 如果密码错误
	var pass string = fmt.Sprint(user["password"])
	err = bcrypt.CompareHashAndPassword([]byte(pass), []byte(userPassword))
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "账号或密码错误，登录失败",
		})
		return
	}
	// 登录成功，获取Token
	token, err := common.ReleaseToken(userTel)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "系统异常，登录失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":  200,
		"token": token,
		"msg":   "登录成功",
	})
}

func Info(ctx *gin.Context) {
	userTel, _ := ctx.Get("user")
	// 读取数据库
	users, err := util.ReadJsons("data/user.json")
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": http.StatusServiceUnavailable,
			"msg":  "用户数据读取出错",
		})
		return
	}
	// 查找用户
	isExist := false
	user := gin.H{}
	for _, item := range users {
		if item["tel"] == userTel {
			isExist = true
			user = item
		}
	}
	fmt.Print(userTel)
	if !isExist {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "用户不存在",
		})
		return
	}
	// 返回用户信息
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"user": user,
	})
}
