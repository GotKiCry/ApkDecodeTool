package sdktool

import (
	"Go_Demo2/utils"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

/**
Jar to smali
*/
func jarDecode(path string) {

}

var rFilePath []string

//ApkDecode 解包APK提取文件
func ApkDecode(apkPath string, output string) {
	apkPath = strings.Replace(apkPath, "\"", "", -1)
	output = strings.Replace(output, "\"", "", -1)

	utils.InitDir(output)

	projectPath := output + "project"
	assets := output + "assets"
	libs := output + "libs"
	res := output + "res"
	smali := output + "smali"

	utils.DecodeApk(apkPath, projectPath+"\\")
	apkFileInfo, _ := os.Stat(apkPath)
	decodeDir := projectPath + "\\" + strings.Replace(apkFileInfo.Name(), ".apk", "", -1) + "\\"
	log.Println("开始整合资源")
	dir, err := ioutil.ReadDir(decodeDir)
	if err != nil {
		log.Println("文件目录打开失败，解包失败")
		log.Fatalln(err)
		return
	}

	for _, file := range dir {
		name := file.Name()
		targetFile := decodeDir + file.Name()
		if err != nil {
			log.Println("目录打开失败 " + decodeDir + file.Name())
			continue
		}
		switch {
		case file.Name() == "assets":
			copyDir(targetFile, assets, "assets")
		case file.Name() == "lib":
			copyDir(targetFile, libs, "libs")
		case file.Name() == "res":
			copyDir(targetFile, res, "res")
		case strings.Index(name, "smali") != -1:
			copyDir(targetFile, smali, "smali")
		default:
			log.Println("略过 " + file.Name())
		}
	}
	log.Println("R文件位置 =====================================================================================")
	fmt.Print("[")
	for _, r := range rFilePath {
		fmt.Printf("\"%s\",", r)
	}
	fmt.Print("]")
	log.Println(" \n请自行判断哪些R是必要，可默认全添加，减少R可提高打包速度")
}

func copyDir(readDir string, writeDir string, ruleName string) error {
	dir, err := ioutil.ReadDir(readDir)
	if err != nil {
		log.Fatalf("file not exist %s", readDir)
		return err
	}

	for _, file := range dir {
		name := file.Name()

		//跳过文件并记录R文件位置
		if isSkip(ruleName, readDir+"\\"+name) {
			if strings.Index(name, "R$") != -1 {
				paths := strings.Split(readDir, "\\")
				var resourcePaths []string
				var pathsIndex = -1
				for index, path := range paths {
					if strings.Index(path, "smali") != -1 {
						//resourcePaths = paths[index+1:]
						pathsIndex = index
						continue
					}
				}

				if pathsIndex == -1 {
					continue
				} else {
					resourcePaths = paths[pathsIndex+1:]
				}
				var resourcePath string
				for index, rPath := range resourcePaths {
					if index == 0 {
						resourcePath = rPath
						continue
					}
					resourcePath = resourcePath + "." + rPath
				}

				for _, p := range rFilePath {
					if resourcePath == p {
						goto Out
					}
				}

				rFilePath = append(rFilePath, resourcePath)
			}
		Out:
			continue
		}

		if file.IsDir() {
			copyDir(readDir+"\\"+name, writeDir+"\\"+name, ruleName)
			continue
		}

		os.MkdirAll(writeDir, os.ModeDir)
		utils.InitFile(writeDir + "\\" + name)
		readFile, _ := os.Open(readDir + "\\" + name)
		writeFile, _ := os.OpenFile(writeDir+"\\"+name, os.O_WRONLY, 0666)
		//log.Printf("CopyFile ===》 %s  ", writeDir+"\\"+name)
		copyFile(readFile, writeFile)
		readFile.Close()
		writeFile.Close()
	}
	return nil
}

func copyFile(readFile *os.File, writeFile *os.File) error {
	buf := make([]byte, 4096)

	for {
		bs, err := readFile.Read(buf)
		if err != nil && err != io.EOF {
			log.Println(err)
			return err
		}

		if bs == 0 {
			//log.Println("复制完毕")
			break
		}

		if _, err := writeFile.Write(buf[:bs]); err != nil {
			log.Println(err)
			return err
		}

	}
	return nil
}
