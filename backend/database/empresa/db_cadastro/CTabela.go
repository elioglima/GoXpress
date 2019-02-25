package db_cadastro

import (
	dbsql "goxpress/backend/database"
	"goxpress/libs"
	"strconv"
)

func New(MyConexao *dbsql.MyConexao) *CTabela {
	c := &CTabela{}
	c.Nome_Classe = "cadastro"
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
	c.Dados.Nome = c.Rows[RecordIndex][2]
	c.Dados.Doc1 = c.Rows[RecordIndex][3]
	c.Dados.Doc2 = c.Rows[RecordIndex][4]
	c.Dados.Doc1_desc = c.Rows[RecordIndex][5]
	c.Dados.Doc2_desc = c.Rows[RecordIndex][6]
	c.Dados.Tipo_pessoa, _ = strconv.Atoi(c.Rows[RecordIndex][7])
}

func (c *CTabela) PesquisaDocs(in_id_empresa int, in_pesquisa string) error {
	c.sSQL = "select * from " + c.Nome_Tabela
	c.sSQL += " where id_empresa = " + strconv.Itoa(in_id_empresa)
	c.sSQL += " and ( "
	c.sSQL += " doc1 = " + libs.Asp(in_pesquisa)
	c.sSQL += " or doc2 = " + libs.Asp(in_pesquisa)
	c.sSQL += " )"

	err := c.Query(c.sSQL)
	if err != nil {
		return err
	}

	if c.RowsCount > 0 {
		c.LoadDados(0)
	}

	return nil
}
