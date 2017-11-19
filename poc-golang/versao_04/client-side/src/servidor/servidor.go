package servidor

import (
	"log"
	"net/http"
	"variaveis"

	"github.com/gorilla/mux"
)

func IniciarServidorWeb() {
	gorillaRoute := mux.NewRouter()
	gorillaRoute.HandleFunc("/", carregarPagina)
	// URL com parametros dinamicos
	gorillaRoute.HandleFunc("/{categoria}", carregarPagina)

	http.HandleFunc("/css/", serveResource)
	http.HandleFunc("/js/", serveResource)

	http.Handle("/", gorillaRoute)
	http.ListenAndServe(":8081", nil)
}

func carregarPagina(w http.ResponseWriter, r *http.Request) {

	atualizarArquivoJSON()
	parametrosURL := mux.Vars(r)
	variaveis.PaginaSelecionada = parametrosURL["categoria"]
	if variaveis.PaginaSelecionada == "" {
		variaveis.PaginaSelecionada = "geral"
	}

	paginaEstatica := paginasEstaticas.Lookup(variaveis.PaginaSelecionada + ".html")
	if paginaEstatica == nil {
		log.Println("NAO ACHOU!!")
		paginaEstatica = paginasEstaticas.Lookup("404.html")
		w.WriteHeader(404)
	}

	//Values to pass into the template
	pagina := variaveis.DefaultContext{}
	pagina.Titulo = variaveis.PaginaSelecionada

	paginaEstatica.Execute(w, pagina)
}
