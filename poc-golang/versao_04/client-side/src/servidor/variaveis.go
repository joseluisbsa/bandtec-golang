package servidor

///home/joseph/github/bandtec-golang/poc-golang/versao_04
var paginaSelecionada string
var temaDaPagina = "bs4"
var public = "/home/joseph/github/bandtec-golang/poc-golang/versao_04/client-side/src/public/"
var pastaPaginas = "/home/joseph/github/bandtec-golang/poc-golang/versao_04/client-side/src/pages"
var pastaTemas = "/home/joseph/github/bandtec-golang/poc-golang/versao_04/client-side/src/themes/"
var localArquivosJSON = "/home/joseph/github/bandtec-golang/poc-golang/versao_04/client-side/src/pages/"
var urlTodasDenuncias = "http://localhost:8080/denuncias/"
var urlTodasDenunciasPorRegiao = "http://localhost:8080/denuncias/0"

type DefaultContext struct {
	Titulo string
}

type DadosDasDenuncias struct {
	ID     string `json:"id,omitempty"`
	Nome   string `json:"nome,omitempty"`
	Total  string `json:"total,omitempty"`
	Regiao string `json:"regiao,omitempty"`
}
