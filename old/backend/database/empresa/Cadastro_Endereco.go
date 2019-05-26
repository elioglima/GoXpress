package empresa

import (
	dbsql "goxpress/backend/database"
	"goxpress/libs"
)

func NewCadastro_Endereco(MyConexao *dbsql.MyConexao) *CCadastro_Endereco {
	c := &CCadastro_Endereco{}
	c.Nome_Classe = libs.InterfaceToName(c)
	c.New(MyConexao)
	return c
}

type CCadastro_EnderecoRow struct {
	Id          int    `db:id`
	Id_empresa  int    `db:id_empresa`
	Id_cadastro int    `db:id_cadastro`
	Endereco    string `db:endereco`
	Numero      int    `db:numero`
	Complemento string `db:complemento`
	Cep         string `db:cep`
	Bairro      string `db:bairro`
	Cidade      string `db:cidade`
	Estado      string `db:estado`
	UF          string `db:uf`
}

type CCadastro_Endereco struct {
	dbsql.DBBase
}
