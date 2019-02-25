package agente

import (
	dbsql "goxpress/backend/database"
	"goxpress/libs"
)

func NewResposta_Ident(MyConexao *dbsql.MyConexao) *CResposta_Ident {
	c := &CResposta_Ident{}
	c.Nome_Classe = libs.InterfaceToName(c)
	c.New(MyConexao)
	return c
}

type CResposta_Ident_Row struct {
	Id            int
	Id_empresa    int
	Id_assistente int
	Descricao     string
	Afirmativa    bool
	Parametros    string
}

type CResposta_Ident struct {
	dbsql.DBBase
}
