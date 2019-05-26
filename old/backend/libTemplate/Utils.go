package libTemplate

import (
	"goxpress/backend/global"
	"html/template"
	"os"
)

var (
	tmpl *template.Template
)

func Parse(name string) (*template.Template, error) {

	fileTmp := global.DirPrivate() + "gohtml/" + name + ".gohtml"

	var err error
	_, err = os.Stat(fileTmp)

	if os.IsNotExist(err) {
		global.Logger.Erro("libTemplate.Parse(string)")
		global.Logger.Erro(err, fileTmp)
		return nil, err
	}

	tmpl = template.Must(template.ParseFiles(fileTmp))
	return tmpl, nil

}
