package ftpClient

/*
	Referecia: http://blog.ralch.com/tutorial/golang-ftp/
*/

import (
	"goxpress/backend/global"

	"github.com/jlaffaye/ftp"
)

var (
	client *ftp.ServerConn
)

func Conectar() error {

	var err error
	global.Logger.Atencao("Iniciando conexão com Servido FTP")

	client, err = ftp.Dial("ftp.maxtime.info:21")
	if err != nil {
		global.Logger.Erro("Erro ao conectar ao FTP")
		return err
	}

	if err := client.Login("maxtime", "AB102030"); err != nil {
		global.Logger.Erro("Erro ao conectar ao FTP")
		return err
	}

	global.Logger.Sucesso("FTP Conectado")
	return nil
}

func MkDir(sDir string) error {

	var err error

	err = client.MakeDir(sDir)
	if err != nil {
		global.Logger.Erro(err)
	}

	return nil
}

func ChangeDir(sDir string) error {

	var err error

	err = client.ChangeDir(sDir)
	if err != nil {
		global.Logger.Erro(err)
	}

	return nil
}

func Fechar() {
	global.Logger.Sucesso("FTP Descoectado")
	client.Quit()
}

func SendFile(FileName string) error {

	// var err error

	// data := bytes.NewReader

	// err = c.Stor("tset", data)
	// if err != nil {
	// 	t.Error(err)
	// }

	return nil
}
