package assistente

import (
	classe "goxpress/backend/Classes"
	"goxpress/backend/database/assistente/db_apresentacao"
	"goxpress/backend/database/assistente/db_assistente"
	"goxpress/backend/database/assistente/db_conversa"
	"goxpress/backend/database/assistente/db_conversa_indicador"
	"goxpress/backend/database/empresa/db_cadastro"
	"goxpress/backend/database/empresa/db_empresa"
	"goxpress/backend/database/empresa/db_registro"
)

var (
	TecWhats     int = 50001
	TecFace      int = 50002
	TecTelegran  int = 50003
	TecInstagran int = 50004
)

type CAgente struct {
	AI                     *classe.CAgenteIdentificador
	Empresa                *db_empresa.CTabela
	Cadastro               *db_cadastro.CTabela
	Assistente             *db_assistente.CTabela
	Apresentacao           *db_apresentacao.CTabela
	Conversa               *db_conversa.CTabela
	Conversa_identificador *db_conversa_indicador.CTabela
	Registro               *db_registro.CTabela
}

func NewCAgente() *CAgente {
	c := &CAgente{}
	return c
}
