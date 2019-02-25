package global

import (
	dbsql "goxpress/backend/database"
	"goxpress/libs"
	"os"
)

const (
	PRIVATE_DIR = "/private/"
	PUBLIC_DIR  = "/public/"
	PORT        = "80"
)

var (
	Logger    *libs.Logs
	MyConexao *dbsql.MyConexao
)

func Load() {
	Logger = libs.NewLogs()
	Logger.Chamada("Load Vars")
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

func ConctarDB() {
	var err error
	MyConexao, err = dbsql.NewMyConexao("mysql.maxtime.info", 3306, "maxtime01", "AB102030", "maxtime01")
	if err != nil {
		Logger.Erro(err)
		return
	}

	Logger.Chamada("Banco de dados conectado.")

}
