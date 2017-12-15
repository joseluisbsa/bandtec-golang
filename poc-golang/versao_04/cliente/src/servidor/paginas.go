package servidor

import (
	"bufio"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// função usada para atualizar os arquivos .html e .json
func atualizarArquivosWeb() {
	paginasEstaticas = carregarHTML()
}

func carregarPagina(w http.ResponseWriter, r *http.Request) {

	atualizarArquivosJSON()
	parametrosURL := mux.Vars(r)
	paginaSelecionada = parametrosURL["categoria"]
	if paginaSelecionada == "" {
		paginaSelecionada = "geral"
	}

	paginaEstatica := paginasEstaticas.Lookup(paginaSelecionada + ".html")
	if paginaEstatica == nil {
		log.Println("NAO ACHOU!!")
		paginaEstatica = paginasEstaticas.Lookup("404.html")
		w.WriteHeader(404)
	}

	pagina := Contexto{}
	pagina.Titulo = paginaSelecionada

	paginaEstatica.Execute(w, pagina)
}

func carregarHTML() *template.Template {
	arquivosHTMLEncontrados := template.New("templates")
	arquivosHTML := new([]string)

	pastaArquivosHTML, erro := os.Open(localArquivosHTMLeJSON)
	verificarErro(erro, "ERRO AO BUSCAR ARQUIVOS HTML E JSON", true)
	defer pastaArquivosHTML.Close()

	todosArquivosHTML, erro := pastaArquivosHTML.Readdir(-1)
	verificarErro(erro, "ERRO AO COLETAR ARQUIVOS DA PASTA HTML E JSON", true)
	for _, arquivo := range todosArquivosHTML {
		//log.Println(arquivo.Name())
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
