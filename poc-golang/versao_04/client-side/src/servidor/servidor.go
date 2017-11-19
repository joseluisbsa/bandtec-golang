package servidor

import (
	"log"
	"net/http"

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

	//Values to pass into the template
	pagina := DefaultContext{}
	pagina.Titulo = paginaSelecionada

	paginaEstatica.Execute(w, pagina)
}
