package controller

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func post(word, s string) string {
	data := make(url.Values)
	data["q"] = []string{word}
	data["from"] = []string{"auto"}
	data["to"] = []string{"auto"}
	data["appid"] = []string{"20211120001004171"}
	data["salt"] = []string{"1435660288"}
	data["sign"] = []string{s}

	res, err := http.PostForm("http://api.fanyi.baidu.com/api/trans/vip/translate", data)
	if err != nil {
		fmt.Println(err.Error())
		return "err"
	}
	defer res.Body.Close()
	fmt.Println("post send success")

	result, _ := ioutil.ReadAll(res.Body)
	return string(result)
}

func Fanyi(ctx *gin.Context) {
	word := ctx.PostForm("word")
	s := fmt.Sprintf("%x", md5.Sum([]byte("20211120001004171"+word+"1435660288"+"Rq5j1ZMu3iE1SY6gFjTg")))
	resp := post(word, s)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": string(resp),
	})
}
