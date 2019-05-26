package agente

import (
	dbsql "goxpress/backend/database"
	"goxpress/libs"
	"strconv"
)

func NewAssistente(MyConexao *dbsql.MyConexao) *CAssistente {
	c := &CAssistente{}
	c.Nome_Classe = libs.InterfaceToName(c)
	c.New(MyConexao)
	return c
}

type CAssistente_Row struct {
	Id         int
	Id_empresa int
	Nome       string
}

type CAssistente struct {
	dbsql.DBBase
	sSQL string
}

func (c *CAssistente) GetID(in_id_empresa, in_id int) error {
	c.sSQL = "select * from " + c.Nome_Tabela
	c.sSQL += " where id_empresa = " + strconv.Itoa(in_id_empresa)
	c.sSQL += " and id = " + strconv.Itoa(in_id)
	return c.Query(c.sSQL)
}

func (c *CAssistente) GetAll(in_id_empresa int) error {
	c.sSQL = "select * from " + c.Nome_Tabela
	c.sSQL += " where id_empresa = " + strconv.Itoa(in_id_empresa)
	return c.Query(c.sSQL)
}
