package agente

/*
	classe define todas as conversas
	relacionamento: id_empresa, id_assistente
*/

import (
	"errors"
	classe "goxpress/backend/Classes"
	dbsql "goxpress/backend/database"
	"goxpress/backend/global"
	"goxpress/libs"
	"strconv"
)

func NewConversa_identificador(MyConexao *dbsql.MyConexao) *CConversa_identificador {
	c := &CConversa_identificador{}
	c.Nome_Classe = libs.InterfaceToName(c)
	c.New(MyConexao)
	return c
}

type CConversa_identificador_Row struct {
	Id                 int
	Id_empresa         int
	Tipo_identificador int
	Identificador      string
	Documento          string
	Fatura             string
}

type CConversa_identificador struct {
	dbsql.DBBase
	sSQL  string
	First *CConversa_identificador_Row
}

func (c *CConversa_identificador) Get(AI *classe.CAgenteIdentificador) (bool, error) {

	if AI.EP == 0 {
		return false, errors.New("Empresa não informado")

	} else if len(AI.Identificador) == 0 {
		return false, errors.New("Identificador não localizado")

	}

	c.sSQL = "select * from " + c.Nome_Tabela
	c.sSQL += " where id_empresa = " + strconv.Itoa(AI.EP)
	c.sSQL += " and tipo_identificador = " + strconv.Itoa(AI.Tecnologia)
	c.sSQL += " and identificador = " + libs.Asp(AI.Identificador)
	c.sSQL += " limit 0,1 "

	err := c.Query(c.sSQL)
	if err != nil {
		return false, err
	}

	if c.RowsCount == 0 {
		return false, nil
	}

	c.First = &CConversa_identificador_Row{}
	c.First.Id, _ = strconv.Atoi(c.Rows[0][0])
	c.First.Id_empresa, _ = strconv.Atoi(c.Rows[0][1])
	c.First.Identificador = c.Rows[0][2]
	c.First.Documento = c.Rows[0][3]
	c.First.Fatura = c.Rows[0][4]

	Sucesso := ((len(c.First.Documento) > 0) || (len(c.First.Fatura) > 0))
	global.Logger.Erro("ok", Sucesso, c.First.Documento, c.First.Fatura)

	return Sucesso, nil
}

func (c *CConversa_identificador) Add(dados CConversa_identificador_Row) error {
	sSQL := "insert into " + c.Nome_Tabela + " "
	sSQL += " (id_empresa, tipo_identificador, identificador, documento, fatura) "
	sSQL += " values ( "
	sSQL += strconv.Itoa(dados.Id_empresa)
	sSQL += "," + strconv.Itoa(dados.Tipo_identificador)
	sSQL += "," + libs.Asp(dados.Identificador)
	sSQL += "," + libs.Asp(dados.Documento)
	sSQL += "," + libs.Asp(dados.Fatura)
	sSQL += " ) "

	err := c.ExecSQL(sSQL)
	if err != nil {
		return errors.New("Erro ao inserir na tabela " + c.Nome_Tabela + " :: " + err.Error())
	}
	return nil
}

func (c *CConversa_identificador) Update(dados CConversa_identificador_Row) error {

	sSQL := "update " + c.Nome_Tabela + " set "
	sSQL += " documento = " + libs.Asp(dados.Documento)
	sSQL += " ,fatura = " + libs.Asp(dados.Fatura)
	sSQL += " where id_empresa = " + strconv.Itoa(dados.Id_empresa)
	sSQL += " and Tipo_identificador = " + strconv.Itoa(dados.Tipo_identificador)
	sSQL += " and identificador = " + libs.Asp(dados.Identificador)

	err := c.ExecSQL(sSQL)
	if err != nil {
		return errors.New("Erro ao update  na tabela " + c.Nome_Tabela + " :: " + err.Error())
	}
	return nil
}
