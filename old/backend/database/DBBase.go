package database

import (
	_ "github.com/ziutek/mymysql/thrsafe"
)

const (
	dbb_tp_int    = 3
	dbb_tp_string = 253
)

type DBBaseColunas struct {
	name string
	tipo byte
}

type DBBase struct {
	Conn        *MyConexao
	Nome_Tabela string
	Nome_Classe string
	ColCount    int
	Cols        []DBBaseColunas

	RowsCount int
	Rows      [][]string
}

func (d *DBBase) New(MyConexao *MyConexao) {
	d.Conn = MyConexao
	Nome_Classe := d.Nome_Classe
	d.Nome_Tabela = "goex_" + Nome_Classe
}

func (d *DBBase) Query(sSQL string) error {

	rows, res, err := d.Conn.DB.Query(sSQL)
	if err != nil {
		return err
	}

	d.ColCount = len(res.Fields())
	for _, field := range res.Fields() {
		BColuna := DBBaseColunas{}
		BColuna.name = string(field.Name)
		BColuna.tipo = field.Type
		d.Cols = append(d.Cols, BColuna)
	}

	d.RowsCount = len(rows)

	for _, row := range rows {
		var valor []string

		for _, field := range res.Fields() {
			colname := res.Map(field.Name)
			valor = append(valor, row.Str(colname))
		}

		d.Rows = append(d.Rows, valor)
	}

	return nil

}

func (d *DBBase) ExecSQL(sSQL string) error {

	stmt, err := d.Conn.DB.Prepare(sSQL)
	if err != nil {
		return err
	}

	_, err = stmt.Run()
	return err

}
