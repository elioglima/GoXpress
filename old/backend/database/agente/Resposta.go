package agente

import (
	dbsql "goxpress/backend/database"
	"goxpress/libs"
)

func NewResposta(MyConexao *dbsql.MyConexao) *CResposta {
	c := &CResposta{}
	c.Nome_Classe = libs.InterfaceToName(c)
	c.New(MyConexao)
	return c
}

type CResposta_Row struct {
	Id            int
	Id_empresa    int
	Id_assistente int
	Id_Resp_Acao  int
	Id_Resp_Ident int
	Parametros    string
	Resposta      string
}

type CResposta struct {
	dbsql.DBBase
}
