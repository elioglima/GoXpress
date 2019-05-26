package db_empresa

import (
	dbsql "goxpress/backend/database"
	"strconv"
)

func New(MyConexao *dbsql.MyConexao) *CTabela {
	c := &CTabela{}
	c.Nome_Classe = "empresa"
	c.New(MyConexao)
	return c
}

type CTabela struct {
	dbsql.DBBase
	sSQL string
}

func (c *CTabela) GetID(in_id int) error {
	c.sSQL = "select * from " + c.Nome_Tabela
	c.sSQL += " where id = " + strconv.Itoa(in_id)
	return c.Query(c.sSQL)
}
