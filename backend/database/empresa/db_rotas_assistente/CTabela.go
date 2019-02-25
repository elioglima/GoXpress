package db_rotas_assistente

import (
	dbsql "goxpress/backend/database"
	"strconv"
)

func New(MyConexao *dbsql.MyConexao) *CTabela {
	c := &CTabela{}
	c.Nome_Classe = "rotas_assistente"
	c.New(MyConexao)
	return c
}

type CTabela struct {
	dbsql.DBBase
	sSQL string
}

func (c *CTabela) GetEtapa(in_id_empresa, in_etapa int) error {
	c.sSQL = "select * from " + c.Nome_Tabela
	c.sSQL += " where id_empresa = " + strconv.Itoa(in_id_empresa)
	c.sSQL += " and ordem = " + strconv.Itoa(in_etapa)
	c.sSQL += " limit 0,1"
	return c.Query(c.sSQL)
}
