package servidor

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"paginas"
	"strings"
	"variaveis"

	"github.com/gorilla/mux"
)

func ServeWeb() {

	gorillaRoute := mux.NewRouter()
	gorillaRoute.HandleFunc("/", serveContent)
	// URL com parametros dinamicos
	gorillaRoute.HandleFunc("/{pageAlias}", serveContent)

	http.HandleFunc("/css/", serveResource)
	http.HandleFunc("/js/", serveResource)

	http.Handle("/", gorillaRoute)
	http.ListenAndServe(":8081", nil)
}

func serveContent(w http.ResponseWriter, r *http.Request) {

	atualizaJSON()
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

func serveResource(w http.ResponseWriter, req *http.Request) {

	//path := "/home/joseph/github/bandtec-golang/poc-golang/versao_03/client-side/src/public/" + themeName + req.URL.Path
	path := "C:/Users/aluno/bandtec-golang/poc-golang/versao_04/client-side/src/public/" + paginas.ThemeName + req.URL.Path
	var contentType string

	if strings.HasSuffix(path, ".css") {
		contentType = "text/css; charset=utf-8"
	} else if strings.HasSuffix(path, ".js") {
		contentType = "application/javascript; charset=utf-8"
	} else {
		contentType = "text/plain; charset=utf-8"
	}

	log.Println(path)
	f, err := os.Open(path)
	if err == nil {
		defer f.Close()
		w.Header().Add("Content-type", contentType)
		br := bufio.NewReader(f)
		br.WriteTo(w)
	} else {
		w.WriteHeader(404)
	}
}
