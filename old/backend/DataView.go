package backend

type SDataView struct {
	Titulo  string
	Versao  string
	UrlHttp string
}

func NewDataView() *SDataView {
	S := &SDataView{}
	S.Titulo = "GoExpress"
	S.Versao = "1.2"
	return S
}
