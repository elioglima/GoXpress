package backend

import (
	"log"
	"os"
)

type CnfArq struct {
	Backend  string
	Frontend string
	Libs     string
	https    string
}

func newCnfArq() *CnfArq {

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	c := &CnfArq{}
	c.Backend = wd + "backend/"
	c.Frontend = wd + "frontend/"
	c.Libs = wd + "lib/"
	return c
}

func (bc *backend) Configs() {

	// configurações
	bc.Arquivos = newCnfArq()

}
