package database

import (
	"errors"
	"strconv"

	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/thrsafe"
)

type MyConexao struct {
	Host    string
	Porta   int
	Usuario string
	Senha   string
	Banco   string

	DB mysql.Conn
}

func NewMyConexao(host string, porta int, usuario, senha, banco string) (MyC *MyConexao, Erro error) {

	if porta == 0 {
		porta = 3306
	}

	MyC = &MyConexao{}
	MyC.DB = mysql.New("tcp", "", host+":"+strconv.Itoa(porta), usuario, senha, banco)
	err := MyC.DB.Connect()
	if err != nil {
		Erro = errors.New("Erro ao efetuar conexão ao banco de dados")
	}

	MyC.Usuario = usuario
	MyC.Senha = senha
	MyC.Host = host
	MyC.Banco = banco

	return

}
