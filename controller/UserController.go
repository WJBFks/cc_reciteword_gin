package controller

import (
	"cc_gatherer_gin/DB"
	"cc_gatherer_gin/common"
	"cc_gatherer_gin/model"
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
	user, err := DB.Users.Query(userTel)
	if err != nil {
		fmt.Print(err)
		ctx.JSON(200, gin.H{
			"code": http.StatusServiceUnavailable,
			"msg":  "用户数据读取出错",
		})
		return
	}

	// 检查手机号唯一性
	if user.Tel == userTel {
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
			"msg":  "系统错误，注册失败1",
		})
	}

	err = DB.Users.Insert(model.User{
		Name:     userName,
		Tel:      userTel,
		Password: string(pass),
		Email:    userEmail,
		Words:    "",
	})
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusServiceUnavailable,
			"msg":  "系统错误，注册失败2",
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
	// 查找用户
	user, err := DB.Users.Query(userTel)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": http.StatusServiceUnavailable,
			"msg":  "账号或密码错误，登录失败",
		})
		return
	}
	// 如果密码错误
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userPassword))
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
	// 查找用户
	user, err := DB.Users.Query(fmt.Sprint(userTel))
	if err != nil {
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

func UsersInfo(ctx *gin.Context) {
	users, err := DB.Users.List()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusServiceUnavailable,
			"msg":  "查找失败",
		})
		return
	}
	// 返回用户信息
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"user": users,
	})
}

func Words(ctx *gin.Context) {
	// 获取参数
	tel := ctx.PostForm("tel")
	words := ctx.PostForm("words")
	err := DB.Users.WordsUpdate(tel, words)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusServiceUnavailable,
			"msg":  "单词表更新失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "单词表更新成功",
	})
}
