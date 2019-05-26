# DB - Assistente 
Documentação do assistente, recebimento de mensagem e interpretador de mensagens de redes sociais.. 

### Assistente  
Rota: /assistente
Metodo: POST

#### request body
{
	"EP":1
	,"AS":1
	,"Tecnologia":5001
	,"Identificador":"11952550331"
	,"Mensagem":"Oi tudo bem."
}

> EP = id da empresa / 
> AS = id do assistente



#### response body
{
    "DataProcessameto": "Feb 6, 2019 at 10:02am (UTC)",
    "Token": null,
    "HttpStatus": 0,
    "HttpResponse": "Autenticando Solicitação"
}
