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

	arquivosHTMLEncontrados.ParseFiles(*arquivosHTML...)
	return arquivosHTMLEncontrados
}

func carregarEstilo(w http.ResponseWriter, req *http.Request) {

	arquivoCSS := localArquivosCSS + req.URL.Path
	log.Println(req.URL.Path)
	var contentType = "text/css; charset=utf-8"

	conteudoDoArquivo, erro := os.Open(arquivoCSS)
	if erro == nil {
		w.Header().Add("Content-type", contentType)
		br := bufio.NewReader(conteudoDoArquivo)
		br.WriteTo(w)
	} else {
		w.WriteHeader(404)
	}
	defer conteudoDoArquivo.Close()
}
