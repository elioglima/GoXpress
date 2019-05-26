package conversa

import "time"

type CRow struct {
	Id            int
	Id_empresa    int
	Id_assistente int
	Responsavel   int
	Identificador string
	Mensagem      string
	Resposta      string
	Data          time.Time
	Inicio        int
}
