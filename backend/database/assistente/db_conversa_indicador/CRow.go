package db_conversa_indicador

/*
	classe define todas as conversas
	relacionamento: id_empresa, id_assistente
*/

type CRow struct {
	Id                   int
	Id_empresa           int
	Identificador        string
	Tecnologia           int
	Documento            string
	Fatura               string
	Etapa                int
	Executado            int
	Msg_repetida_cliente int
}
