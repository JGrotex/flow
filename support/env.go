package support

import (
	"github.com/project-flogo/core/engine"
	"os"
)

const (
	UserName   = "FLOGO_APP_USERNAME"
	HostName   = "FLOGO_HOST_NAME"
	AppName    = "FLOGO_APP_NAME"
	AppVersion = "FLOGO_APP_VERSION"
)

var username, hostName, appName, appVersion string

func GetUserName() string {
	if len(username) > 0 {
		return username
	}
	username = os.Getenv(UserName)
	if len(username) > 0 {
		return username
	}
	return "flogo"
}

func GetHostId() string {
	if len(hostName) > 0 {
		return hostName
	}
	hostName = os.Getenv(HostName)
	if len(hostName) > 0 {
		return hostName
	}
	h, _ := os.Hostname()
	return h
}

func GetAppName() string {
	if len(appName) > 0 {
		return appName
	}
	appName = os.Getenv(AppName)
	if len(appName) <= 0 {
		return engine.GetAppName()
	}
	return appName
}

func GetAppVerison() string {
	if len(appVersion) > 0 {
		return appVersion
	}
	appVersion = os.Getenv(AppVersion)
	if len(appVersion) <= 0 {
		return engine.GetAppVersion()
	}
	return appVersion
}
