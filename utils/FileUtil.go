package utils

import (
	"fmt"
	"io"
	"log"
	"os"
)

func WriteMsgToFile(filePath string, content string) {
	openFile, _ := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	io.WriteString(openFile, content+"\r\n")
}

// InitFile 初始化文件
func InitFile(filePath string) {
	os.Remove(filePath)
	_, err := os.Create(filePath)
	if err != nil {
		log.Println(err)
	}
	//fmt.Printf("清理文件 %s \r\n", filePath)
}

// InitDir 初始化文件夹
func InitDir(filePath string) {
	fmt.Printf("清理目录 %s \r\n", filePath)
	os.RemoveAll(filePath)
	os.MkdirAll(filePath, os.ModeDir)
}
