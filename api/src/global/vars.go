package global

import (
	"os"

	"../libs"
)

const (
	PRIVATE_DIR = "/private/"
	PUBLIC_DIR  = "/public/"
	PORT        = "80"
)

var (
	Logger *libs.Logs
)

func Load() {
	Logger = libs.NewLogs()
	// Logger.Chamada("Load Vars")
	LoadConfigs()
}

func DirPublic() string {
	return PUBLIC_DIR
}

func DirPrivate() string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}

	return dir + PRIVATE_DIR
}
