package main

import (
	"Go_Demo2/utils"
)

const sdkProjectPath = "E:\\_WorkCode\\supersdkchannelas"

//const apkPath = "\"E:\\_WorkCode\\supersdkchannelas\\app\\build\\intermediates\\apk\\debug\\app-debug.apk\""

const apkPath = "\"E:\\_WorkCode\\supersdkchannelas\\app\\build\\outputs\\apk\\release\\app-release.apk\""

func main() {
	f1 := "D:\\QP_共享文件夹\\三国如龙传-威狐新-1.0.3.1-20220707142035"
	utils.BackCodeApk(f1, f1+"\\output", 0 == 0)
	//utils.DecodeApk(f1, f1+"\\..\\")
	//sdktool.ApkDecode(apkPath, "E:\\_Tools\\jar转smali\\out\\")
	//utils.SignApkV1(f1, f1)
	//utils.GetFileAssets(assets)
	//decodeAllApk()
	//backcodeAllApk()
	//mt := "com\\\\qipa"
	//path := "E:\\_Tools\\jar转smali\\out\\project\\app-debug\\smali_classes9\\com\\qipa"
	//fmt.Println(regexp.MatchString(mt, path))
	//ca-app-pub-5689749980636351/1073386940

	//path := "D:\\QiDianFile\\20220630\\付佳薇"
	//dir, _ := ioutil.ReadDir(path)
	//for _, o := range dir {
	//	f := path + "\\" + o.Name()
	//	//utils.DecodeApk(f, f+"\\..\\")
	//	utils.BackCodeApk(f, f+"\\output", 0 == 0)
	//}

}
