package empresa

import (
	dbsql "goxpress/backend/database"
	"goxpress/libs"
	"strconv"
)

func NewEmpresa(MyConexao *dbsql.MyConexao) *CEmpresa {
	c := &CEmpresa{}
	c.Nome_Classe = libs.InterfaceToName(c)
	c.New(MyConexao)
	return c
}

type CEmpresa_Row struct {
	Id           int    `db:id`
	Nome         string `db:nome`
	Doc1         string `db:doc1`
	Doc2         string `db:doc2`
	NomeFantasia string `db:nome_fantasia`
}

type CEmpresa struct {
	dbsql.DBBase
	sSQL string
}

func (c *CEmpresa) GetID(in_id int) error {
	c.sSQL = "select * from " + c.Nome_Tabela
	c.sSQL += " where id = " + strconv.Itoa(in_id)
	return c.Query(c.sSQL)
}
