package db_usuario

import (
	dbsql "goxpress/backend/database"
)

func New(MyConexao *dbsql.MyConexao) *CTabela {
	c := &CTabela{}
	c.Nome_Classe = "usuario"
	c.New(MyConexao)
	return c
}

type CTabela struct {
	dbsql.DBBase
	Row string
}
