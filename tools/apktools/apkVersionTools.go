package apktools

import (
	"github.com/lunny/axmlParser"
)


func GetApkInfo(path string)(appName, appVesionCode, appVersionName string){
	listener := new(axmlParser.AppNameListener)
	axmlParser.ParseApk(path, listener)

	appName = listener.ActivityName
	appVesionCode = listener.VersionCode
	appVersionName = listener.VersionName

	return
}
