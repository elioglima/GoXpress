package empresa

import (
	dbsql "goxpress/backend/database"
	"goxpress/libs"
)

func NewCadastro_Contato(MyConexao *dbsql.MyConexao) *CCadastro_Contato {
	c := &CCadastro_Contato{}
	c.Nome_Classe = libs.InterfaceToName(c)
	c.New(MyConexao)
	return c
}

type CCadastro_Contato_Row struct {
	Id                       int    `db:id`
	Id_empresa               int    `db:id_empresa`
	Id_cadastro              int    `db:id_cadastro`
	Id_cadastro_contato_tipo int    `db:id_cadastro_contato_tipo`
	valor                    string `db:valor`
}

type CCadastro_Contato struct {
	dbsql.DBBase
}
