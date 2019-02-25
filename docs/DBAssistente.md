# DB - Assistente 
Documentação do assistente

### Assistente Resposta Ação - Adicionar 
Rota: /db/assistente/resposta_acao/add 
Metodo: POST


#### request body
{
	"id_empresa":1
	,"id_assistente":1
	,"descricao":"Teste"
	,"ativo":true
}

#### response body
{
    "DataProcessameto": "Feb 4, 2019 at 4:18pm (UTC)",
    "Token": null,
    "HttpStatus": 201,
    "HttpResponse": "O registro ja existe na base de dados"
}


### Assistente Resposta Ação - Editar 
Rota: /db/assistente/resposta_acao/edd 
Metodo: POST


#### request body
{
	"id":1
	,"id_empresa":1
	,"id_assistente":1
	,"descricao":"Descrição 02"
	,"ativo":false
}


#### response body
{
    "DataProcessameto": "Feb 5, 2019 at 10:19am (UTC)",
    "Token": null,
    "HttpStatus": 201,
    "HttpResponse": "Dados atualizado com sucesso"
}