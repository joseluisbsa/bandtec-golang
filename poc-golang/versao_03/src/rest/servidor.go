package rest

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func IniciarServidor() {

	router := mux.NewRouter()

	router.HandleFunc("/denuncias/", pegarTodasCategorias).Methods("GET")   // JSON com todas as categorias
	router.HandleFunc("/denuncias/{uri}", pegarUmaCategoria).Methods("GET") // devolve apenas uma categoria
	router.HandleFunc("/denuncias/", gravarNovaDenuncia).Methods("POST")    // adiciona nova denuncia

	log.Fatal(http.ListenAndServe(":8080", router)) // Server na porta 8080 [ localhost:8080 ]
}
