package db_cadastro_contato_tipo

import (
	dbsql "goxpress/backend/database"
)

func New(MyConexao *dbsql.MyConexao) *CTabela {
	c := &CTabela{}
	c.Nome_Classe = "cadastro_contato_tipo"
	c.New(MyConexao)
	return c
}

type CTabela struct {
	dbsql.DBBase
}
