package db_cadastro_contato

import (
	dbsql "goxpress/backend/database"
)

func New(MyConexao *dbsql.MyConexao) *Ctabela {
	c := &Ctabela{}
	c.Nome_Classe = "cadastro_contato"
	c.New(MyConexao)
	return c
}

type Ctabela struct {
	dbsql.DBBase
}
