package main

import "Go_Demo2/utils"

// "Go_Demo2/utils"
// "io/ioutil"

const sdkProjectPath = "E:\\_WorkCode\\supersdkchannelas"

//const apkPath = "\"E:\\_WorkCode\\supersdkchannelas\\app\\build\\intermediates\\apk\\debug\\app-debug.apk\""

const apkPath = "\"E:\\_WorkCode\\supersdkchannelas\\app\\build\\outputs\\apk\\release\\app-release.apk\""

const f1 = "\"E:\\_打包\\打包目录\\苍之纪元\\苍之纪元-GM商城更新工具-13.2-20220823151842.apk\""

const t = "E:\\_WorkCode\\PrivateDialog\\app\\build\\outputs\\apk\\release\\app-release-unsigned.apk"

func main() {

	//sdktool.GenNewStyleableXml()

	//utils.BackCodeApk(f1, f1+"\\output", 1 == 1)
	utils.DecodeApk(f1, f1+"\\..\\")
	//sdktool.ApkDecode(apkPath, "E:\\_Tools\\jar转smali\\out\\")
	//utils.SignApkV1(f1, f1)
	//utils.GetFileAssets(assets)
	//decodeAllApk()
	// backcodeAllApk()Destination directory
	//mt := "com\\\\qipa"
	//path := "E:\\_Tools\\jar转smali\\out\\project\\app-debug\\smali_classes9\\com\\qipa"
	//fmt.Println(regexp.MatchString(mt, path))
	//ca-app-pub-5689749980636351/1073386940

	// utils.RunCmd("adb","install",)

	// path := "D:\\QiDianFile\\20220801\\朱银洁"
	// dir, _ := ioutil.ReadDir(path)
	// for _, o := range dir {
	// 	f := path + "\\" + o.Name()
	// 	// utils.DecodeApk(f, f+"\\..\\")
	// 	utils.BackCodeApk(f, f+"\\..\\output", 0 == 0)
	// }

}
