package db_conversa_indicador

/*
	classe define todas as conversas
	relacionamento: id_empresa, id_assistente
*/

import (
	"errors"
	classe "goxpress/backend/Classes"
	dbsql "goxpress/backend/database"
	"goxpress/libs"
	"strconv"
)

func New(MyConexao *dbsql.MyConexao) *CTabela {
	c := &CTabela{}
	c.Nome_Classe = "conversa_identificador"
	c.New(MyConexao)
	return c
}

type CTabela struct {
	dbsql.DBBase
	sSQL  string
	Dados *CRow
}

func (c *CTabela) LoadDados(RecordIndex int) {
	c.Dados = &CRow{}
	c.Dados.Id, _ = strconv.Atoi(c.Rows[RecordIndex][0])
	c.Dados.Id_empresa, _ = strconv.Atoi(c.Rows[RecordIndex][1])
	c.Dados.Tecnologia, _ = strconv.Atoi(c.Rows[RecordIndex][2])
	c.Dados.Identificador = c.Rows[RecordIndex][3]
	c.Dados.Documento = c.Rows[RecordIndex][4]
	c.Dados.Fatura = c.Rows[RecordIndex][5]
	c.Dados.Etapa, _ = strconv.Atoi(c.Rows[RecordIndex][6])

}
func (c *CTabela) Get(AI *classe.CAgenteIdentificador) (bool, error) {

	if AI.EP == 0 {
		return false, errors.New("Empresa não informado")

	} else if len(AI.Identificador) == 0 {
		return false, errors.New("Identificador não localizado")

	}

	c.sSQL = "select * from " + c.Nome_Tabela
	c.sSQL += " where id_empresa = " + strconv.Itoa(AI.EP)
	c.sSQL += " and tipo_identificador = " + strconv.Itoa(AI.Tecnologia)
	c.sSQL += " and identificador = " + libs.Asp(AI.Identificador)
	c.sSQL += " order by id desc "
	c.sSQL += " limit 0,1 "

	err := c.Query(c.sSQL)
	if err != nil {
		return false, err
	}

	if c.RowsCount == 0 {
		return false, nil
	}

	c.LoadDados(0)
	Sucesso := ((len(c.Dados.Documento) > 0) || (len(c.Dados.Fatura) > 0))
	return Sucesso, nil
}

func (c *CTabela) IncRepeticaoMsg(AI *classe.CAgenteIdentificador) error {
	sSQL := " update " + c.Nome_Tabela + " "
	sSQL += " msg_repetida_cliente = msg_repetida_cliente+1"
	sSQL += " where id_empresa = " + strconv.Itoa(AI.EP)
	sSQL += " and tipo_identificador = " + strconv.Itoa(AI.Tecnologia)
	sSQL += " and identificador = " + libs.Asp(AI.Identificador)
	err := c.ExecSQL(sSQL)
	if err != nil {
		return errors.New("Erro ao update  na tabela " + c.Nome_Tabela + " :: " + err.Error())
	}
	return nil
}

func (c *CTabela) Update(dados *CRow) error {

	sSQL := "update " + c.Nome_Tabela + " set "
	sSQL += " documento = " + libs.Asp(dados.Documento)
	sSQL += " ,fatura = " + libs.Asp(dados.Fatura)
	sSQL += " ,msg_repetida_cliente = " + strconv.Itoa(dados.Msg_repetida_cliente)
	sSQL += " where id = " + strconv.Itoa(dados.Id)
	sSQL += " and  id_empresa = " + strconv.Itoa(dados.Id_empresa)
	sSQL += " and Tipo_identificador = " + strconv.Itoa(dados.Tecnologia)
	sSQL += " and identificador = " + libs.Asp(dados.Identificador)

	err := c.ExecSQL(sSQL)
	if err != nil {
		return errors.New("Erro ao update  na tabela " + c.Nome_Tabela + " :: " + err.Error())
	}
	return nil
}

func (c *CTabela) addMsgRepetidaCliente(dados *CRow) error {
	sSQL := "update " + c.Nome_Tabela + " set "
	sSQL += " msg_repetida_cliente = msg_repetida_cliente+1"
	sSQL += " where id = " + strconv.Itoa(dados.Id)
	sSQL += " and  id_empresa = " + strconv.Itoa(dados.Id_empresa)
	sSQL += " and Tipo_identificador = " + strconv.Itoa(dados.Tecnologia)
	sSQL += " and identificador = " + libs.Asp(dados.Identificador)

	err := c.ExecSQL(sSQL)
	if err != nil {
		return errors.New("Erro ao update  na tabela " + c.Nome_Tabela + " :: " + err.Error())
	}
	return nil
}

func (c *CTabela) AddMsgRepetidaCliente(dados *CRow) error {
	identificador := New(c.Conn)
	return identificador.addMsgRepetidaCliente(dados)
}

func (c *CTabela) Add(dados *CRow) error {
	sSQL := "insert into " + c.Nome_Tabela + " "
	sSQL += " (id_empresa, tipo_identificador, identificador, documento, fatura, etapa) "
	sSQL += " values ( "
	sSQL += strconv.Itoa(dados.Id_empresa)
	sSQL += "," + strconv.Itoa(dados.Tecnologia)
	sSQL += "," + libs.Asp(dados.Identificador)
	sSQL += "," + libs.Asp(dados.Documento)
	sSQL += "," + libs.Asp(dados.Fatura)
	sSQL += "," + strconv.Itoa(dados.Etapa)
	sSQL += " ) "

	err := c.ExecSQL(sSQL)
	if err != nil {
		return errors.New("Erro ao inserir na tabela " + c.Nome_Tabela + " :: " + err.Error())
	}

	return nil
}
