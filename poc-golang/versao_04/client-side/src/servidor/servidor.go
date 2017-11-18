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
	gorillaRoute.HandleFunc("/", serveContent)
	// URL com parametros dinamicos
	gorillaRoute.HandleFunc("/{pageAlias}", serveContent)

	http.HandleFunc("/css/", paginas.ServeResource)
	http.HandleFunc("/js/", paginas.ServeResource)

	http.Handle("/", gorillaRoute)
	http.ListenAndServe(":8081", nil)
}

func serveContent(w http.ResponseWriter, r *http.Request) {

	atualizarJSON()
	urlParams := mux.Vars(r)
	variaveis.PageAlias = urlParams["pageAlias"]
	if variaveis.PageAlias == "" {
		variaveis.PageAlias = "geral"
	}

	staticPage := paginas.StaticPages.Lookup(variaveis.PageAlias + ".html")
	if staticPage == nil {
		log.Println("NAO ACHOU!!")
		staticPage = paginas.StaticPages.Lookup("404.html")
		w.WriteHeader(404)
	}

	//Values to pass into the template
	context := variaveis.DefaultContext{}
	context.Title = variaveis.PageAlias

	staticPage.Execute(w, context)
}
