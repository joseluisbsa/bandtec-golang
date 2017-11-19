package servidor

import (
	"bufio"
	"html/template"
	"log"
	"net/http"
	"os"
)

var paginasEstaticas = populateStaticPages()

// função usada para atualizar os arquivos .html e .json
func atualizarArquivosWeb() {
	paginasEstaticas = populateStaticPages()
}

func populateStaticPages() *template.Template {
	resultado := template.New("templates")
	templatePaths := new([]string)

	pastaTemplate, erro := os.Open(pastaPaginas)
	if erro != nil {
		log.Println(erro)
	}
	defer pastaTemplate.Close()

	arquivosPastaTemplate, erro := pastaTemplate.Readdir(-1)
	if erro != nil {
		log.Println(erro)
	}

	for _, pathinfo := range arquivosPastaTemplate {
		log.Println(pathinfo.Name())
		*templatePaths = append(*templatePaths, pastaPaginas+"/"+pathinfo.Name())
	}

	basePath := pastaTemas + temaDaPagina
	pastaTemplate, erro = os.Open(basePath)
	if erro != nil {
		log.Println(erro)
	}
	defer pastaTemplate.Close()
	arquivosPastaTemplate, erro = pastaTemplate.Readdir(-1)
	if erro != nil {
		log.Println(erro)
	}
	for _, pathinfo := range arquivosPastaTemplate {
		log.Println(pathinfo.Name())
		*templatePaths = append(*templatePaths, basePath+"/"+pathinfo.Name())
	}

	resultado.ParseFiles(*templatePaths...)
	return resultado
}

func carregarEstilo(w http.ResponseWriter, req *http.Request) {

	arquivosCSS := pastaEstilo + req.URL.Path
	log.Println(req.URL.Path)
	var contentType = "text/css; charset=utf-8"

	conteudoPasta, err := os.Open(arquivosCSS)
	if err == nil {
		defer conteudoPasta.Close()
		w.Header().Add("Content-type", contentType)
		br := bufio.NewReader(conteudoPasta)
		br.WriteTo(w)
	} else {
		w.WriteHeader(404)
	}
}
