package global

var (
	private_PrjTitulo      string
	private_HttpHome       string
	private_HttpSolicitado string
)

func LoadConfigs() {
	Logger.Chamada("LoadConfigs")
	private_PrjTitulo = "GoXpress"
}

func PrjTitulo() string {
	return private_PrjTitulo
}

func SetHttpHome(valor string) {
	private_HttpHome = valor
}

func HttpHome() string {
	return private_HttpHome
}

func SetHttpSolicitado(valor string) {
	private_HttpSolicitado = valor
}

func HttpSolicitado() string {
	return private_HttpSolicitado
}
