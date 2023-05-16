package main

import "Go_Demo2/utils"

const sdkProjectPath = "E:\\_WorkCode\\supersdkchannelas"

//const apkPath = "\"E:\\_WorkCode\\supersdkchannelas\\app\\build\\intermediates\\apk\\debug\\app-debug.apk\""

const apkPath = "\"E:\\_WorkCode\\supersdkchannelas\\app\\build\\outputs\\apk\\release\\app-release.apk\""

const f1 = "D:\\Code\\AndroidProject\\ApktoolDemo\\app\\build\\intermediates\\apk\\debug\\app-debug.apk"

const test = "C:\\Users\\liuzh\\Documents\\Tencent Files\\3008479817\\FileRecv\\乌龙院之活宝传奇_小七_20230428150925.apk"

const t = "E:\\_WorkCode\\PrivateDialog\\app\\build\\outputs\\apk\\release\\app-release-unsigned.apk"

func main() {

	// sdktool.GenNewStyleableXml("E:\\_Tools\\jar转smali\\out\\project\\app-release")

	// utils.BackCodeApk(f1, f1+"\\output", 1 == 1)

	utils.DecodeApk(f1, f1+"\\..\\")
	// utils.DecodeApk(test, test+"\\..\\")

	// sdktool.ApkDecode(apkPath, "E:\\_Tools\\jar转smali\\out\\")
	// utils.SignApkV1(f1, f1)
	//utils.GetFileAssets(assets)
	//decodeAllApk()
	// backcodeAllApk()Destination directory
	//mt := "com\\\\qipa"
	//path := "E:\\_Tools\\jar转smali\\out\\project\\app-debug\\smali_classes9\\com\\qipa"
	//fmt.Println(regexp.MatchString(mt, path))
	//ca-app-pub-5689749980636351/1073386940

	// utils.RunCmd("adb","install",)

	// path := "C:\\Users\\liuzh\\Documents\\Tencent Files\\3008479817\\FileRecv\\"

	// dir, _ := os.ReadDir(path)
	// for _, o := range dir {
	// 	f := path + "\\" + o.Name()
	// 	// utils.DecodeApk(f, f+"\\..\\")
	// 	utils.BackCodeApk(f, f+"\\..\\output", 0 == 0)
	// }

}
