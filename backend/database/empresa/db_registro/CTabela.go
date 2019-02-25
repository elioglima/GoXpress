package db_registro

import (
	dbsql "goxpress/backend/database"
	"goxpress/backend/global"
	"goxpress/libs"
	"strconv"
)

func New(MyConexao *dbsql.MyConexao) *CTabela {
	c := &CTabela{}
	c.Nome_Classe = "registro"
	c.New(MyConexao)
	return c
}

type CTabela struct {
	dbsql.DBBase
	sSQL  string
	Dados *CRow
}

func (c *CTabela) LoadDados(RecordIndex int) {
	c.Dados = &CRow{}
	c.Dados.Id, _ = strconv.Atoi(c.Rows[RecordIndex][0])
	c.Dados.Id_empresa, _ = strconv.Atoi(c.Rows[RecordIndex][1])
	c.Dados.Id_cadastro, _ = strconv.Atoi(c.Rows[RecordIndex][2])
	c.Dados.Nome = c.Rows[RecordIndex][3]
	c.Dados.Doc1 = c.Rows[RecordIndex][4]
	c.Dados.Doc2 = c.Rows[RecordIndex][5]
	c.Dados.NrFatura = c.Rows[RecordIndex][6]

	valor_debito, err := strconv.ParseFloat(c.Rows[RecordIndex][7], 64)
	if err != nil {
		global.Logger.Erro(err)
		return
	}

	c.Dados.Valor_debito = valor_debito

}

func (c *CTabela) Pesquisa(in_id_empresa int, valor string) error {
	c.sSQL = "select * from " + c.Nome_Tabela
	c.sSQL += " where id_empresa = " + strconv.Itoa(in_id_empresa)
	c.sSQL += " and ("
	c.sSQL += " nrfatura = " + libs.Asp(valor)
	c.sSQL += " or doc1 = " + libs.Asp(valor)
	c.sSQL += " or doc2 = " + libs.Asp(valor)
	c.sSQL += ")"
	c.sSQL += " limit 0,1"

	err := c.Query(c.sSQL)
	if err != nil {
		return err
	}

	if c.RowsCount > 0 {
		c.LoadDados(0)
	}

	return nil
}
