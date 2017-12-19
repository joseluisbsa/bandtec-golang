package servidor

var paginaSelecionada string
var temaDaPagina = "bs4"

var localArquivosHTMLeJSON = "/home/joseph/github/bandtec-golang/poc-golang/versao_04/cliente/src/paginas"

//var localArquivosHTMLeJSON = "D:/github/bandtec-golang/poc-golang/versao_04/client-side/src/paginas"

var localArquivosCSS = "/home/joseph/github/bandtec-golang/poc-golang/versao_04/cliente/src"

//var localArquivosCSS = "D:/github/bandtec-golang/poc-golang/versao_04/client-side/src"
var urlTodasDenuncias = "http://12.0.0.92:8080/denuncias/"
var urlTodasDenunciasPorRegiao = "http://12.0.0.92:8080/denuncias/0"

var paginasEstaticas = carregarHTML()

type Contexto struct {
	Titulo string
}

type DadosDasDenuncias struct {
	ID     string `json:"id,omitempty"`
	Nome   string `json:"nome,omitempty"`
	Total  string `json:"total,omitempty"`
	Regiao string `json:"regiao,omitempty"`
}
