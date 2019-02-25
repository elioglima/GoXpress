package db_cadastro_endereco

type CRow struct {
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
