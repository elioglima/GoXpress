package empresa

import (
	dbsql "goxpress/backend/database"
	"goxpress/libs"
)

func NewCadastro(MyConexao *dbsql.MyConexao) *CCadastro {
	c := &CCadastro{}
	c.Nome_Classe = libs.InterfaceToName(c)
	c.New(MyConexao)
	return c
}

type CCadastro_Row struct {
	Id         int    `db:"id"`
	Id_empresa int    `db:"id_empresa"`
	Nome       string `db:"nome"`
	Doc1       string `db:"doc1"`
	Doc2       string `db:"doc2"`
}

type CCadastro struct {
	dbsql.DBBase
}
