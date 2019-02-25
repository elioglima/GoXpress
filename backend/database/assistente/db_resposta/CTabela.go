package db_resposta

import (
	dbsql "goxpress/backend/database"
	"goxpress/libs"
)

func New(MyConexao *dbsql.MyConexao) *CTabela {
	c := &CTabela{}
	c.Nome_Classe = libs.InterfaceToName(c)
	c.New(MyConexao)
	return c
}

type CTabela struct {
	dbsql.DBBase
}
