package db_cadastro

type CRow struct {
	Id          int    `db:"id"`
	Id_empresa  int    `db:"id_empresa"`
	Nome        string `db:"nome"`
	Doc1        string `db:"doc1"`
	Doc2        string `db:"doc2"`
	Doc1_desc   string `db:"doc1_desc"`
	Doc2_desc   string `db:"doc2_desc"`
	Tipo_pessoa int    `db:"tipo_pessoa"` // NOT NULL
}
