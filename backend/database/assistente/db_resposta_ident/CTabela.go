package db_resposta_ident

import (
	dbsql "goxpress/backend/database"
)

func New(MyConexao *dbsql.MyConexao) *CTabela {
	c := &CTabela{}
	c.Nome_Classe = "resposta_ident"
	c.New(MyConexao)
	return c
}

type CTabela struct {
	dbsql.DBBase
}
