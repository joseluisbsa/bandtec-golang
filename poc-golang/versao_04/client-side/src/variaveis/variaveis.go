package variaveis

var PageAlias string
var TemaDaPagina = "bs4"
var Public = "C:/Users/aluno/bandtec-golang/poc-golang/versao_04/client-side/src/public/"
var PastaPaginas = "C:/Users/aluno/bandtec-golang/poc-golang/versao_04/client-side/src/pages"
var PastaTemas = "C:/Users/aluno/bandtec-golang/poc-golang/versao_04/client-side/src/themes/"
var ArquivosJSON = "C:/Users/aluno/bandtec-golang/poc-golang/versao_04/client-side/src/pages/"
var TodasDenuncias = "http://localhost:8080/denuncias/"
var TodasDenunciasPorRegiao = "http://localhost:8080/denuncias/0"

type DefaultContext struct {
	Title string
}

type CategoriaFull struct {
	ID    string `json:"id,omitempty"`
	Nome  string `json:"nome,omitempty"`
	Total string `json:"total,omitempty"`
}

type CategoriaEach struct {
	ID     string `json:"id,omitempty"`
	Nome   string `json:"nome,omitempty"`
	Regiao string `json:"regiao,omitempty"`
	Total  string `json:"total,omitempty"`
}
