package apresentacao

/*

	CApresentacao_Row
		Inicio
			1 = Apresentação inicial
			2 = Apresentação quando ja tem conversa e ja foi confirmado o documentocpf, cnpj ou fatura
			3 = Apresentação quando ja tem conversa porem não foi identificado o numero de cpf ou fatura

*/

import (
	"errors"
	dbsql "goxpress/backend/database"
	"goxpress/backend/global"
	"strconv"
)

func New(MyConexao *dbsql.MyConexao) *CTabela {
	c := &CTabela{}
	c.Nome_Classe = "apresentacao"
	c.New(MyConexao)
	return c
}

type CTabela struct {
	dbsql.DBBase
	sSQL string
}

func (c *CTabela) GetAll(in_id_assistente, in_id_empresa int) error {
	c.sSQL = "select * from " + c.Nome_Tabela
	c.sSQL += " where id_empresa = " + strconv.Itoa(in_id_empresa)
	c.sSQL += " and id_assistente = " + strconv.Itoa(in_id_assistente)
	return c.Query(c.sSQL)
}

func (c *CTabela) GetID(in_id_assistente, in_id_empresa int) error {
	c.sSQL = "select parametros from " + c.Nome_Tabela
	c.sSQL += " where id_empresa = " + strconv.Itoa(in_id_empresa)
	c.sSQL += " and id_assistente = " + strconv.Itoa(in_id_assistente)
	c.sSQL += " limit 0,1 "
	return c.Query(c.sSQL)
}

func (c *CTabela) GetInicio(in_id_assistente, in_id_empresa int) error {
	c.sSQL = "select parametros from " + c.Nome_Tabela
	c.sSQL += " where id_empresa = " + strconv.Itoa(in_id_empresa)
	c.sSQL += " and id_assistente = " + strconv.Itoa(in_id_assistente)
	c.sSQL += " and inicio = 1 "
	c.sSQL += " limit 0,1 "

	err := c.Query(c.sSQL)

	if err != nil {
		global.Logger.Erro(err)
		return err
	}

	// validação do asistente virtual
	if c.RowsCount == 0 {
		return errors.New("Nenhuma apresentação localizada")
	}

	return nil
}

func (c *CTabela) GetContinua(in_id_assistente, in_id_empresa int) error {
	c.sSQL = "select parametros from " + c.Nome_Tabela
	c.sSQL += " where id_empresa = " + strconv.Itoa(in_id_empresa)
	c.sSQL += " and id_assistente = " + strconv.Itoa(in_id_assistente)
	c.sSQL += " and inicio = 2 "
	c.sSQL += " limit 0,1 "

	err := c.Query(c.sSQL)

	if err != nil {
		global.Logger.Erro(err)
		return err
	}

	// validação do asistente virtual
	if c.RowsCount == 0 {
		return errors.New("Nenhuma apresentação localizada")
	}

	return nil

}

func (c *CTabela) GetContinuaSemIdent(in_id_assistente, in_id_empresa int) error {
	c.sSQL = "select parametros from " + c.Nome_Tabela
	c.sSQL += " where id_empresa = " + strconv.Itoa(in_id_empresa)
	c.sSQL += " and id_assistente = " + strconv.Itoa(in_id_assistente)
	c.sSQL += " and inicio = 3 "
	c.sSQL += " limit 0,1 "

	err := c.Query(c.sSQL)

	if err != nil {
		global.Logger.Erro(err)
		return err
	}

	// validação do asistente virtual
	if c.RowsCount == 0 {
		return errors.New("Nenhuma apresentação localizada")
	}

	return nil

}
