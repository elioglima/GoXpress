package agente

import (
	dbsql "goxpress/backend/database"
	"goxpress/libs"
	"strconv"
)

func NewApresentacao(MyConexao *dbsql.MyConexao) *CApresentacao {
	c := &CApresentacao{}
	c.Nome_Classe = libs.InterfaceToName(c)
	c.New(MyConexao)
	return c
}

type CApresentacao_Row struct {
	Id            int
	Id_empresa    int
	Id_assistente int
	parametros    string
}

type CApresentacao struct {
	dbsql.DBBase
	sSQL string
}

func (c *CApresentacao) GetAll(in_id_assistente, in_id_empresa int) error {
	c.sSQL = "select * from " + c.Nome_Tabela
	c.sSQL += " where id_empresa = " + strconv.Itoa(in_id_empresa)
	c.sSQL += " and id_assistente = " + strconv.Itoa(in_id_assistente)
	return c.Query(c.sSQL)
}

func (c *CApresentacao) GetID(in_id_assistente, in_id_empresa int) error {
	c.sSQL = "select parametros from " + c.Nome_Tabela
	c.sSQL += " where id_empresa = " + strconv.Itoa(in_id_empresa)
	c.sSQL += " and id_assistente = " + strconv.Itoa(in_id_assistente)
	c.sSQL += " limit 0,1 "
	return c.Query(c.sSQL)
}
