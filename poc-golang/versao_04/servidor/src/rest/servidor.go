package rest

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// IniciarServidor cria os endereços das requisições
func IniciarServidor() {

	router := mux.NewRouter()

	router.HandleFunc("/denuncias/", pegarTodasDenuncias).Methods("GET")
	router.HandleFunc("/denuncias/{uri}", pegarDenunciasPorRegiao).Methods("GET")
	router.HandleFunc("/denuncias/", gravarNovaDenuncia).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router)) // Server na porta 8080 [ localhost:8080 ]
}
