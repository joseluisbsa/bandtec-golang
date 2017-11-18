package servidor

import (
	"log"
	"net/http"
	"paginas"
	"variaveis"

	"github.com/gorilla/mux"
)

func ServeWeb() {
	gorillaRoute := mux.NewRouter()
	gorillaRoute.HandleFunc("/", conteudoDaPagina)
	// URL com parametros dinamicos
	gorillaRoute.HandleFunc("/{pageAlias}", conteudoDaPagina)

	http.HandleFunc("/css/", paginas.ServeResource)
	http.HandleFunc("/js/", paginas.ServeResource)

	http.Handle("/", gorillaRoute)
	http.ListenAndServe(":8081", nil)
}

func conteudoDaPagina(w http.ResponseWriter, r *http.Request) {

	atualizarJSON()
	urlParams := mux.Vars(r)
	variaveis.PageAlias = urlParams["pageAlias"]
	if variaveis.PageAlias == "" {
		variaveis.PageAlias = "geral"
	}

	paginaEstatica := paginas.PaginasEstaticas.Lookup(variaveis.PageAlias + ".html")
	if paginaEstatica == nil {
		log.Println("NAO ACHOU!!")
		paginaEstatica = paginas.PaginasEstaticas.Lookup("404.html")
		w.WriteHeader(404)
	}

	//Values to pass into the template
	pagina := variaveis.DefaultContext{}
	pagina.Titulo = variaveis.PageAlias

	paginaEstatica.Execute(w, pagina)
}
