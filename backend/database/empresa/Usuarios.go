package empresa

import (
	dbsql "goxpress/backend/database"
	"goxpress/libs"
)

func NewUsuario(MyConexao *dbsql.MyConexao) *CUsuario {
	c := &CUsuario{}
	c.Nome_Classe = libs.InterfaceToName(c)
	c.New(MyConexao)
	return c
}

type CUsuario_Row struct {
	Id         int    `db:"id"`
	Id_empresa int    `db:"codigo_empresa"`
	Nome       string `db:"nome"`
}

type CUsuario struct {
	dbsql.DBBase
	Row string
}
