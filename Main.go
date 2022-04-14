package main

import "Go_Demo2/utils"

const sdkProjectPath = "E:\\_WorkCode\\supersdkchannelas"

const apkPath = "\"E:\\_WorkCode\\supersdkchannelas\\app\\build\\intermediates\\apk\\debug\\app-debug.apk\""

func main() {
	f1 := "D:\\QiDianFile\\3008479817\\朱银洁\\20220414\\文明曙光-乐澄-4.0.2-20220410221709_20220414_朱银洁_2850909328"
	utils.BackCodeApk(f1, f1+"\\output", true)
	//utils.DecodeApk(f1, f1+"\\..\\")
	//utils.SignApkV1(f1, f1)
	//utils.GetFileAssets(assets)
	//decodeAllApk()
	//backcodeAllApk()
	//sdktool.ApkDecode(apkPath, "E:\\_Tools\\jar转smali\\out\\")
}
