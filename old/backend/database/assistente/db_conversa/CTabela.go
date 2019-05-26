package db_conversa

/*
	classe define todas as conversas
	relacionamento: id_empresa, id_assistente
*/

import (
	"errors"
	dbsql "goxpress/backend/database"
	"goxpress/libs"
	"strconv"
)

func New(MyConexao *dbsql.MyConexao) *CTabela {
	c := &CTabela{}
	c.Nome_Classe = "conversa"
	c.New(MyConexao)
	return c
}

type CTabela struct {
	dbsql.DBBase
	sSQL string
}

func (c *CTabela) Get(dados *CRow) (bool, error) {

	if dados.Id_empresa == 0 {
		return false, errors.New("Empresa não informado")

	} else if dados.Id_assistente == 0 {
		return false, errors.New("Assistente não informado")

	} else if len(dados.Identificador) == 0 {
		return false, errors.New("Identificador não localizado")

	}

	c.sSQL = "select * from " + c.Nome_Tabela
	c.sSQL += " where id_empresa = " + strconv.Itoa(dados.Id_empresa)
	c.sSQL += " and id_assistente = " + strconv.Itoa(dados.Id_assistente)
	c.sSQL += " and identificador = " + libs.Asp(dados.Identificador)
	c.sSQL += " and responsavel = 1" // verifica a ultima mensagem do usuario
	c.sSQL += " order by id desc "
	c.sSQL += " limit 0,1 "

	err := c.Query(c.sSQL)
	if err != nil {
		return false, err
	}

	if c.RowsCount == 0 {
		return false, nil
	}

	return true, nil
}

func (c *CTabela) Gravar(dados *CRow) error {
	sSQL := "insert into " + c.Nome_Tabela + " "
	sSQL += " (id_empresa, id_assistente, identificador, mensagem, Responsavel, Data, Inicio, Resposta) "
	sSQL += " values ( "
	sSQL += strconv.Itoa(dados.Id_empresa)
	sSQL += "," + strconv.Itoa(dados.Id_assistente)
	sSQL += "," + libs.Asp(dados.Identificador)
	sSQL += "," + libs.Asp(dados.Mensagem)
	sSQL += ", " + strconv.Itoa(dados.Responsavel)
	sSQL += ", Now()"
	sSQL += "," + strconv.Itoa(dados.Inicio)
	sSQL += "," + libs.Asp(dados.Resposta)
	sSQL += " ) "

	err := c.ExecSQL(sSQL)
	if err != nil {
		return errors.New("Erro ao gravar conversa :: " + err.Error())
	}
	return nil
}
