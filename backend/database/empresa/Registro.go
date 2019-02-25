package empresa

import (
	dbsql "goxpress/backend/database"
	"goxpress/libs"
	"strconv"
)

func NewRegistro(MyConexao *dbsql.MyConexao) *CRegistro {
	c := &CRegistro{}
	c.Nome_Classe = libs.InterfaceToName(c)
	c.New(MyConexao)
	return c
}

type CRegistro_Row struct {
	Id          int
	Id_empresa  int
	Id_cadastro int
	Nome        string
	Doc1        string
	Doc2        string
	NrFatura    string
}

type CRegistro struct {
	dbsql.DBBase
	sSQL string
}

func (c *CRegistro) Pesquisa(in_id_empresa int, valor string) error {
	c.sSQL = "select 1 from " + c.Nome_Tabela
	c.sSQL += " where id_empresa = " + strconv.Itoa(in_id_empresa)
	c.sSQL += " and ("
	c.sSQL += " nrfatura = " + libs.Asp(valor)
	c.sSQL += " or doc1 = " + libs.Asp(valor)
	c.sSQL += " or doc2 = " + libs.Asp(valor)
	c.sSQL += ")"
	c.sSQL += " limit 0,1"
	return c.Query(c.sSQL)
}
