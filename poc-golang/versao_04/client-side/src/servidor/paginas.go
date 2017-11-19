package servidor

import (
	"bufio"
	"html/template"
	"log"
	"net/http"
	"os"
)

var paginasEstaticas = carregarHTML()

// função usada para atualizar os arquivos .html e .json
func atualizarArquivosWeb() {
	paginasEstaticas = carregarHTML()
}

func carregarHTML() *template.Template {
	arquivosHTMLEncontrados := template.New("templates")
	arquivosHTML := new([]string)

	pastaArquivosHTML, erro := os.Open(localArquivosHTMLeJSON)
	verificarErro(erro)

	defer pastaArquivosHTML.Close()

	todosArquivosHTML, erro := pastaArquivosHTML.Readdir(-1)
	verificarErro(erro)

	for _, arquivo := range todosArquivosHTML {
		log.Println(arquivo.Name())
		*arquivosHTML = append(*arquivosHTML, localArquivosHTMLeJSON+"/"+arquivo.Name())
	}

	pastaArquivosHTML, erro = os.Open(pastaTemas)
	verificarErro(erro)
	defer pastaArquivosHTML.Close()

	todosArquivosHTML, erro = pastaArquivosHTML.Readdir(-1)
	verificarErro(erro)

	for _, arquivo := range todosArquivosHTML {
		log.Println(arquivo.Name())
		*arquivosHTML = append(*arquivosHTML, pastaTemas+"/"+arquivo.Name())
	}

	arquivosHTMLEncontrados.ParseFiles(*arquivosHTML...)
	return arquivosHTMLEncontrados
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
