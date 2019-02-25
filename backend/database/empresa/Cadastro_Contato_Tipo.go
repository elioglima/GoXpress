package empresa

import (
	dbsql "goxpress/backend/database"
	"goxpress/libs"
)

func NewCadastro_Contato_Tipo(MyConexao *dbsql.MyConexao) *CCadastro_Contato_Tipo {
	c := &CCadastro_Contato_Tipo{}
	c.Nome_Classe = libs.InterfaceToName(c)
	c.New(MyConexao)
	return c
}

type CCadastro_Contato_Tipo_Row struct {
	Id         int    `db:id`
	Id_empresa int    `db:id_empresa`
	Descricao  string `db:descricao`
	Ativo      bool   `db:ativo`
}

type CCadastro_Contato_Tipo struct {
	dbsql.DBBase
}
