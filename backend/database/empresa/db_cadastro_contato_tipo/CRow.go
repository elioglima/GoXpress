package db_cadastro_contato_tipo

type CRow struct {
	Id         int    `db:id`
	Id_empresa int    `db:id_empresa`
	Descricao  string `db:descricao`
	Ativo      bool   `db:ativo`
}
