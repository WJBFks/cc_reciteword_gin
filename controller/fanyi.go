package controller

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func get(url string) string {

	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}

	return result.String()
}

func post(url string, data interface{}, contentType string) string {

	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return string(result)
}

func Fanyi(ctx *gin.Context) {
	word := ctx.PostForm("word")
	s := fmt.Sprintf("%x", md5.Sum([]byte("20211120001004171"+word+"1435660288"+"Rq5j1ZMu3iE1SY6gFjTg")))
	resp := get("http://api.fanyi.baidu.com/api/trans/vip/translate?q=" + word +
		"&from=auto&to=auto&appid=20211120001004171&salt=1435660288&sign=" + s)

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": string(resp),
	})
}
