package util

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func SaveByte(data []byte, filename string) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return fmt.Errorf("打开文件失败")
	}
	defer f.Close()
	f.Write(data)
	f.Close()
	return nil
}

func SaveJson(data gin.H, filename string) error {
	b, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("序列化错误")
	}
	return SaveByte(b, filename)
}

func SaveJsons(data []gin.H, filename string) error {
	b, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("序列化错误")
	}
	return SaveByte(b, filename)
}

func ReadByte(filename string) ([]byte, error) {
	f, err := os.OpenFile(filename, os.O_RDONLY, 0777)
	if err != nil {
		return []byte(""), fmt.Errorf("打开文件失败")
	}
	b := make([]byte, 1024*2)
	n, err := f.Read(b)
	if err != nil {
		return []byte(""), fmt.Errorf("文件读取失败")
	}
	return b[:n], nil
}

func ReadJson(filename string) (gin.H, error) {
	b, err := ReadByte(filename)
	if err != nil {
		return gin.H{}, err
	}
	data := gin.H{}
	if fmt.Sprint(b) == "[]" {
		return gin.H{}, nil
	}
	err = json.Unmarshal(b, &data)
	if err != nil {
		return gin.H{}, fmt.Errorf("JSON解析失败")
	}
	return data, nil
}

func ReadJsons(filename string) ([]gin.H, error) {
	b, err := ReadByte(filename)
	if err != nil {
		return []gin.H{}, err
	}
	data := []gin.H{}
	if fmt.Sprint(b) == "[]" {
		return []gin.H{}, nil
	}
	err = json.Unmarshal(b, &data)
	if err != nil {
		return []gin.H{}, fmt.Errorf("JSON解析失败")
	}
	return data, nil
}
