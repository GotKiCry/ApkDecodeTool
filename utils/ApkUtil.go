package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

const assets = "D:\\新建文件夹\\output\\织女星计划-顺网-1.1.9.1-20220116112922\\assets"

func DecodeApk(apkPath string, outputPath string) {
	apkPath = strings.Replace(apkPath, "\"", "", -1)
	outputPath = strings.Replace(outputPath, "\"", "", -1)
	fmt.Println("-------------------开始DecodeApk-------------------")
	apkToolPath := GetApkToolPath()
	if apkToolPath == "" {
		fmt.Println("反编译APK终止")
		errorEnd()
		return
	}

	apkFile, getFileError := os.Stat(apkPath)
	if getFileError != nil {
		log.Fatalf("APK文件不存在 %s\n", getFileError)
		errorEnd()
		return

	}

	ext := path.Ext(apkPath)
	decodeOutputPath := outputPath + strings.Replace(apkFile.Name(), ext, "", -1)
	err := RunCmd("java", "-jar", apkToolPath, "d", apkPath, "--only-main-classes", "-out", decodeOutputPath)
	if err == io.EOF {
		fmt.Println("-------------------DecodeApk结束-------------------")
		fmt.Println("输出地址：", decodeOutputPath)
	} else {
		errorEnd()
		return
	}
}

func BackCodeApk(codePath string, outputApkPath string, needSign bool) {
	codePath = strings.Replace(codePath, "\"", "", -1)
	outputApkPath = strings.Replace(outputApkPath, "\"", "", -1)

	fmt.Println("-------------------开始BackCodeApk-------------------")
	apkToolPath := GetApkToolPath()
	if apkToolPath == "" {
		fmt.Println("回编译APK终止")
		errorEnd()
		return
	}

	apkFile, getFileError := os.Stat(codePath)
	if getFileError != nil {
		log.Fatalf("待编译文件夹不存在 %s\n", getFileError)
		errorEnd()
		return
	}
	apkOutputPath := outputApkPath + "\\" + apkFile.Name() + ".apk"
	error := RunCmd("java", "-jar", apkToolPath, "b", codePath, "--only-main-classes", "-out", apkOutputPath)
	if error == io.EOF {
		fmt.Println("-------------------BackCodeApk结束-------------------")
		if needSign {
			SignApkV1(apkOutputPath, apkOutputPath)
		}
		fmt.Println("输出地址：", apkOutputPath)
	} else {
		errorEnd()
		return
	}
}

func GetApkToolPath() (apktoolPath string) {
	projectPath, _ := os.Getwd()
	apkToolPath := projectPath + "\\assets\\apktool.jar"
	_, apkToolNotFound := os.Stat(projectPath + "\\assets\\apktool.jar")
	if apkToolNotFound != nil {
		path, _ := os.Getwd()
		fmt.Println("test = ", path)
		log.Fatalf("ApkTool.jar文件未找到，请检查assets目录  %s\n", apkToolNotFound)
		return ""
	}
	return apkToolPath
}

func errorEnd() {
	fmt.Println("-------------------执行出错-------------------")
}

func GetFileAssets(filePath string) {
	InitFile(".\\log.txt")
	getFile(filePath, ".\\log.txt")
}

func getFile(filePath string, localFilePath string) {
	dir, fileNotExits := ioutil.ReadDir(filePath)
	if fileNotExits != nil {
		log.Fatalf("assets不存在 %s", fileNotExits)
		return
	}

	for _, file := range dir {
		if file.IsDir() {
			getFile(filePath+"\\"+file.Name(), localFilePath)
		} else {
			path := filePath + "\\" + file.Name()
			newPath := strings.Replace(path, assets, "- assets", -1)
			content := strings.Replace(newPath, "\\", "/", -1)
			if localFilePath != "" {
				WriteMsgToFile(localFilePath, content)
			}
			fmt.Println(content)
		}
	}

}
