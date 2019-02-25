package assistente

import (
	"fmt"
	classe "goxpress/backend/Classes"
	"goxpress/backend/database/assistente/db_apresentacao"
	"goxpress/backend/database/assistente/db_assistente"
	"goxpress/backend/database/assistente/db_conversa"
	"goxpress/backend/database/assistente/db_conversa_indicador"
	"goxpress/backend/database/empresa/db_cadastro"
	"goxpress/backend/database/empresa/db_empresa"
	"goxpress/backend/database/empresa/db_registro"
	"goxpress/backend/database/empresa/db_rotas_assistente"
	"goxpress/backend/global"
	"goxpress/libs"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func (c *CAgente) Start(AI *classe.CAgenteIdentificador) {

	var (
		IsUltimaConversa      bool
		sucesso_identificador bool
		sucesso_conversa      bool
		sMsgResposta          string
		mensagem_antiga       string
		resposta_antiga       string
	)

	// iniciando bibliotecas
	c.Empresa = db_empresa.New(global.MyConexao)
	c.Apresentacao = db_apresentacao.New(global.MyConexao)
	c.Conversa = db_conversa.New(global.MyConexao)
	c.Cadastro = db_cadastro.New(global.MyConexao)
	c.Registro = db_registro.New(global.MyConexao)
	c.Assistente = db_assistente.New(global.MyConexao)

	c.AI = AI

	// criar regra para localizar empresa apartir do identificado
	// exp: telefone celular
	err := c.Empresa.GetID(c.AI.EP)
	if err != nil {
		global.Logger.Erro(err)
		return
	}

	nome_empresa := c.Empresa.Rows[0][4]
	if len(nome_empresa) == 0 {
		global.Logger.Erro("Nome da empresa não configurado.")
		return
	}

	err = c.Assistente.GetID(c.AI.AS, c.AI.EP)
	if err != nil {
		global.Logger.Erro(err)
		return
	}

	// validação do asistente virtual
	if c.Assistente.RowsCount == 0 {
		global.Logger.Erro("Assistente virtual nnão cadastrado.")
		return
	}

	nome_assistente := c.Assistente.Rows[0][2]
	if len(nome_empresa) == 0 {
		global.Logger.Erro("Nome do assistente não configurado.")
		return
	}

	conversa := &db_conversa.CRow{}
	conversa.Id_assistente = AI.AS
	conversa.Id_empresa = AI.EP
	conversa.Identificador = AI.Identificador
	conversa.Mensagem = AI.Mensagem
	conversa.Responsavel = 1
	conversa.Inicio = 0

	// checa a existencia da ultima conversa
	sucesso_conversa, err = c.Conversa.Get(conversa)
	if err != nil {
		global.Logger.Erro(err)
		return
	}

	if sucesso_conversa {

		mensagem_antiga = c.Conversa.Rows[0][4]
		resposta_antiga = c.Conversa.Rows[0][5]
		IsUltimaConversa = (mensagem_antiga == conversa.Mensagem)

		if IsUltimaConversa {
			global.Logger.Erro("Ultima mensagem localizada é a mesma :: " + conversa.Mensagem)
		}

		// verifica se ja existe um documento relacionado ao identificador
		conversa_identificador := db_conversa_indicador.New(global.MyConexao)
		sucesso_identificador, err = conversa_identificador.Get(c.AI)
		if err != nil {
			global.Logger.Erro(err)
			return
		}

		if sucesso_identificador {

			sDoc := conversa_identificador.Rows[0][4]
			sFatura := conversa_identificador.Rows[0][5]

			DocLoadlizado := ((len(sDoc) > 0) && (len(sFatura) > 0))

			// Conversa Localizada - Identificador confirmação de parametros
			global.Logger.Cor(color.FgHiYellow, "Identificador localizado")
			if DocLoadlizado {
				global.Logger.Cor(color.FgHiYellow, "Documento Localizado")
			}

			iEtapa, _ := strconv.Atoi(conversa_identificador.Rows[0][6])
			iExecutado, _ := strconv.Atoi(conversa_identificador.Rows[0][7])
			// if (!DocLoadlizado) && (iEtapa != 1) {
			// 	iEtapa = 1
			// } else if iExecutado == 1 {
			// 	iEtapa++
			// 	iExecutado = 0
			// }

			global.Logger.Cor(color.FgHiYellow, "Etapa", iEtapa, "Executado", iExecutado)

			RetornoVerificacao := c.ProcessamentoEtapa(conversa, &iEtapa, &iExecutado)
			if RetornoVerificacao.DocumentoEncontrado && !DocLoadlizado {
				iEtapa++

				// atualiza o identificador e etapa
				id, _ := strconv.Atoi(conversa_identificador.Rows[0][0])
				identificador := &db_conversa_indicador.CRow{}
				identificador.Id = id
				identificador.Id_empresa = AI.EP
				identificador.Tecnologia = AI.Tecnologia
				identificador.Identificador = AI.Identificador
				identificador.Etapa = iEtapa
				identificador.Documento = RetornoVerificacao.NDoc
				identificador.Fatura = RetornoVerificacao.NFatura
				identificador.Msg_repetida_cliente = 0

				conversa_identificador := db_conversa_indicador.New(global.MyConexao)
				err = conversa_identificador.Update(identificador)
				if err != nil {
					global.Logger.Erro(err)
				}

			} else {

				// etapa localizada
				id, _ := strconv.Atoi(conversa_identificador.Rows[0][0])
				identificador := &db_conversa_indicador.CRow{}
				identificador.Id = id
				identificador.Id_empresa = AI.EP
				identificador.Tecnologia = AI.Tecnologia
				identificador.Identificador = AI.Identificador

				err = conversa_identificador.AddMsgRepetidaCliente(identificador)
				if err != nil {
					global.Logger.Erro(err)
				}

				rotas_assistente := db_rotas_assistente.New(global.MyConexao)
				err := rotas_assistente.GetEtapa(AI.EP, iEtapa)
				if err != nil {
					global.Logger.Erro(err)
					return
				}

				sMsgResposta = rotas_assistente.Rows[0][3]
				chk_dados := strings.Contains(sMsgResposta, "{cadastro.nome}")
				if chk_dados {

					var (
						docFormatado string = ""
					)

					cadastro := db_cadastro.New(global.MyConexao)
					err := cadastro.PesquisaDocs(AI.EP, conversa_identificador.Dados.Documento)
					if err != nil {
						global.Logger.Erro(err)
						return
					}

					sMsgResposta = strings.Replace(sMsgResposta, "{cadastro.nome}", cadastro.Dados.Nome, -1)
					if strings.Contains(sMsgResposta, "{cadastro.doc1_desc}") {
						sMsgResposta = strings.Replace(sMsgResposta, "{cadastro.doc1_desc}", cadastro.Dados.Doc1_desc, -1)
					}

					if strings.Contains(sMsgResposta, "{cadastro.doc1}") {
						docFormatado = libs.FormataCPF(cadastro.Dados.Doc1)
						sMsgResposta = strings.Replace(sMsgResposta, "{cadastro.doc1}", docFormatado, -1)
					}

					if strings.Contains(sMsgResposta, "{INFOFATURA}") {

						var (
							textoFatura string = ""
						)

						registro := db_registro.New(global.MyConexao)
						err = registro.Pesquisa(AI.EP, conversa_identificador.Dados.Fatura)
						if err != nil {
							global.Logger.Erro(err)
							return
						}

						textoFatura = "\nNº da Fatura: " + registro.Dados.NrFatura
						var valorFormatado = libs.RenderFloat("##.###,##", registro.Dados.Valor_debito)
						textoFatura += "\nValor Total do Débito: R$ " + fmt.Sprintf("%v", valorFormatado)
						sMsgResposta = strings.Replace(sMsgResposta, "{INFOFATURA}", textoFatura, -1)
					}

				}
			}

		} else {

			// etata para receptar respostas e mensagens vinda do cliente
			global.Logger.Cor(color.FgHiYellow, "Identificador nao localizado sem parametros")

			RetornoVerificacao := c.ProcessamentoPrimeiroContatoMensagem(conversa)

			identificador := &db_conversa_indicador.CRow{}
			identificador.Id_empresa = AI.EP
			identificador.Tecnologia = AI.Tecnologia
			identificador.Identificador = AI.Identificador
			identificador.Etapa = 1

			if !RetornoVerificacao.DocumentoEncontrado {
				err = c.Apresentacao.GetContinuaSemIdent(c.AI.AS, c.AI.EP)
				if err != nil {
					global.Logger.Erro(err)
					return
				}

				conversa_identificador := db_conversa_indicador.New(global.MyConexao)
				err = conversa_identificador.Add(identificador)
				if err != nil {
					global.Logger.Erro(err)
				}

				sMsgResposta = c.Apresentacao.Rows[0][0]

			} else if RetornoVerificacao.DocumentoEncontrado {

				identificador.Documento = RetornoVerificacao.NDoc
				identificador.Fatura = RetornoVerificacao.NFatura

				conversa_identificador := db_conversa_indicador.New(global.MyConexao)
				err = conversa_identificador.Add(identificador)
				if err != nil {
					global.Logger.Erro(err)
				}

				rotas_assistente := db_rotas_assistente.New(global.MyConexao)
				err := rotas_assistente.GetEtapa(AI.EP, identificador.Etapa+1)
				if err != nil {
					global.Logger.Erro(err)
					return
				}

				if rotas_assistente.RowsCount > 0 {
					sMsgResposta = rotas_assistente.Rows[0][3]

				}

			}
		}

	} else {

		global.Logger.Cor(color.FgHiYellow, "Grava Primeira conversa")
		conversa.Inicio = 1

		err = c.Apresentacao.GetInicio(c.AI.AS, c.AI.EP)
		if err != nil {
			global.Logger.Erro(err)
			return
		}

		sMsgResposta = c.Apresentacao.Rows[0][0]
		if len(nome_empresa) == 0 {
			global.Logger.Erro("Mensagem de apresentação não configurado.")
			return
		}

	}

	// identificadores de tags
	sMsgResposta = strings.Replace(sMsgResposta, "{assistente.nome}", nome_assistente, -1)
	sMsgResposta = strings.Replace(sMsgResposta, "{empresa.nome}", nome_empresa, -1)

	conversa.Resposta = sMsgResposta

	if (mensagem_antiga != conversa.Mensagem) || (resposta_antiga != conversa.Resposta) {
		err = c.Conversa.Gravar(conversa)
		if err != nil {
			global.Logger.Erro(err)
			return
		}
	}

	// global.Logger.Erro(AI)
	global.Logger.Cor(color.FgCyan, "Usuario :: "+conversa.Mensagem)
	global.Logger.Cor(color.FgHiBlue, "Bot :: "+conversa.Resposta)
}

func (c *CAgente) Close() {

}

type cProcessRetorno struct {
	NFatura             string
	NDoc                string
	DocumentoEncontrado bool
}

func (c *CAgente) PesquisaFatura(in_id_empresa int, in_fatura string) (bool, error) {
	err := c.Registro.Pesquisa(in_id_empresa, in_fatura)
	if err != nil {
		global.Logger.Erro(err)
		return false, err
	}

	if c.Registro.RowsCount == 0 {
		return false, nil
	}

	return true, nil
}

func (c *CAgente) PesquisaDocs(in_id_empresa int, in_doc string) (bool, error) {
	err := c.Cadastro.PesquisaDocs(in_id_empresa, in_doc)
	if err != nil {
		global.Logger.Erro(err)
		return false, err
	}

	if c.Cadastro.RowsCount == 0 {
		return false, nil
	}

	return true, nil
}

func (c *CAgente) ProcessamentoPrimeiroContatoMensagem(conversa *db_conversa.CRow) (Retorno *cProcessRetorno) {
	return c.PesquisaDocumento(conversa)

}

func (c *CAgente) ProcessamentoEtapa(conversa *db_conversa.CRow, iEtapa *int, iExecutado *int) (Retorno *cProcessRetorno) {
	return c.PesquisaDocumento(conversa)
}

func (c *CAgente) PesquisaDocumento(conversa *db_conversa.CRow) (Retorno *cProcessRetorno) {

	var (
		sucesso bool  = false
		err     error = nil
	)

	Retorno = &cProcessRetorno{}

	/*
		processamento de resposta
		tipo de verificação
		1º verificar se é afirmativa, negativa, interrogativa

	*/

	// verificar
	Frase := strings.Split(conversa.Mensagem, " ")
	if len(Frase) == 1 {

		// pesquisa doc1, doc2
		// na tabela cadastro
		err := c.Cadastro.PesquisaDocs(conversa.Id_empresa, conversa.Mensagem)
		if err != nil {
			global.Logger.Erro(err)
			return Retorno
		}

		Retorno.DocumentoEncontrado = (c.Cadastro.RowsCount > 0)
		if Retorno.DocumentoEncontrado {
			Retorno.NDoc = conversa.Mensagem
			return Retorno
		}

		sucesso, err = c.PesquisaFatura(conversa.Id_empresa, conversa.Mensagem)
		if err != nil {
			global.Logger.Erro(err)
			return Retorno
		}

		if sucesso {
			Retorno.NDoc = c.Registro.Rows[0][4]
			Retorno.NFatura = c.Registro.Rows[0][6]
			Retorno.DocumentoEncontrado = true
			return Retorno

		} else {

			sucesso, err = c.PesquisaDocs(conversa.Id_empresa, conversa.Mensagem)
			if err != nil {
				global.Logger.Erro(err)
				return Retorno
			}

			if sucesso {
				Retorno.NDoc = c.Cadastro.Rows[0][3]
				Retorno.NFatura = ""
				Retorno.DocumentoEncontrado = true
				return Retorno
			}

		}

	} else {

		Retorno.DocumentoEncontrado = false

		// quando for apenas uma palavra
		for _, FraseValue := range Frase {

			sucesso, err = c.PesquisaFatura(conversa.Id_empresa, string(FraseValue))
			if err != nil {
				global.Logger.Erro(err)
				return Retorno
			}

			if sucesso {
				Retorno.NDoc = c.Registro.Rows[0][4]
				Retorno.NFatura = c.Registro.Rows[0][6]
				Retorno.DocumentoEncontrado = true
				return Retorno

			} else {

				sucesso, err = c.PesquisaDocs(conversa.Id_empresa, string(FraseValue))
				if err != nil {
					global.Logger.Erro(err)
					return Retorno
				}

				if sucesso {
					Retorno.NDoc = c.Cadastro.Rows[0][3]
					Retorno.NFatura = ""
					Retorno.DocumentoEncontrado = true
					return Retorno
				}

			}

		}

	}

	Retorno.NDoc = ""
	Retorno.NFatura = ""
	Retorno.DocumentoEncontrado = false
	return Retorno
}
