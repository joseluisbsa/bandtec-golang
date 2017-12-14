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
	http.HandleFunc("/css/", carregarEstilo)
	http.Handle("/", gorillaRoute)

	log.Fatal(http.ListenAndServe(":8081", nil))
}
