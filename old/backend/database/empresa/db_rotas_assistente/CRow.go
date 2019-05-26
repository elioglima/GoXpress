package db_rotas_assistente

type CRow struct {
	Id                int
	Id_empresa        int
	ordem             int
	pergunta          string
	parametros        string
	resposta_negativa string
	resposta_positiva string
}
