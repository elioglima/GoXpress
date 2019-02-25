package db_empresa

type CRow struct {
	Id           int    `db:id`
	Nome         string `db:nome`
	Doc1         string `db:doc1`
	Doc2         string `db:doc2`
	NomeFantasia string `db:nome_fantasia`
}
