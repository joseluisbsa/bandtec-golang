package variaveis

///home/joseph/github/bandtec-golang/poc-golang/versao_04
var PaginaSelecionada string
var TemaDaPagina = "bs4"
var Public = "/home/joseph/github/bandtec-golang/poc-golang/versao_04/client-side/src/public/"
var PastaPaginas = "/home/joseph/github/bandtec-golang/poc-golang/versao_04/client-side/src/pages"
var PastaTemas = "/home/joseph/github/bandtec-golang/poc-golang/versao_04/client-side/src/themes/"
var ArquivosJSON = "/home/joseph/github/bandtec-golang/poc-golang/versao_04/client-side/src/pages/"
var URLTodasDenuncias = "http://localhost:8080/denuncias/"
var URLTodasDenunciasPorRegiao = "http://localhost:8080/denuncias/0"

type DefaultContext struct {
	Titulo string
}

type DadosDasDenuncias struct {
	ID     string `json:"id,omitempty"`
	Nome   string `json:"nome,omitempty"`
	Total  string `json:"total,omitempty"`
	Regiao string `json:"regiao,omitempty"`
}
