package db_cadastro_contato

type CRow struct {
	Id                       int    `db:id`
	Id_empresa               int    `db:id_empresa`
	Id_cadastro              int    `db:id_cadastro`
	Id_cadastro_contato_tipo int    `db:id_cadastro_contato_tipo`
	valor                    string `db:valor`
}
