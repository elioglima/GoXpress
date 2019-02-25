package db_usuario

type CRow struct {
	Id         int    `db:"id"`
	Id_empresa int    `db:"codigo_empresa"`
	Nome       string `db:"nome"`
}
