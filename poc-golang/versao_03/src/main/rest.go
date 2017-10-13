// Todo codigo em Go precisa iniciar com este package
package main

// Importa todas as bibliotecas que serão utilizadas no código
import (

	// bibliotecas Nativas do Golang
	"bd"
	"denuncia"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	// Pode ser usado bibliotecas não nativas
	// /gorilla/mux para o servidor HTTP
	// /denisenkom/go-mssqldb para a comunicação com Banco SQLServer
	"github.com/gorilla/mux"
)

// Adicona mais uma denuncia
func gravarNovaDenuncia(w http.ResponseWriter, req *http.Request) {
	// modelo que deve enviado
	// {"categoria":"4","localidade":"2"}
	log.Printf("Post mais uma Nova Denuncia")
	var NovaD denuncia.NovaDenuncia

	// grava em 'novaD' os dados enviados
	erro := json.NewDecoder(req.Body).Decode(&NovaD)
	if erro != nil {
		log.Println("erro em ao gravar em novaD: ", erro.Error())
	}
	// imprime no terminal os valores recebidos
	fmt.Println(NovaD)
	bd.GravarNovaDenuncia(&NovaD)
	bd.AtualizarTodasDenuncias()
}

// função para enviar apenas uma categoria com o total por regiao
func pegarUmaCategoria(w http.ResponseWriter, req *http.Request) {
	// OBSERVAÇÂO: comentarios de como funciona esta na 'func GetUsuario'
	log.Printf("Get uma Categoria")
	parametros := mux.Vars(req)
	var categoriaEncontrada []denuncia.DadosDasDenuncias
	for _, item := range bd.DenunciasPorCategoria {
		if strings.ToLower(item.Nome) == strings.ToLower(parametros["uri"]) {
			categoriaEncontrada = append(categoriaEncontrada, item)
		}
	}
	//json.NewEncoder(w).Encode(denunciasPorCategoria)
	json.NewEncoder(w).Encode(categoriaEncontrada)
}

// envia os dados das categorias via GET
func pegarTodasCategorias(w http.ResponseWriter, req *http.Request) {
	log.Printf("Get categorias")
	json.NewEncoder(w).Encode(bd.Denuncias)
}

func main() {

	bd.AtualizarTodasDenuncias()
	router := mux.NewRouter()

	router.HandleFunc("/denuncias/", pegarTodasCategorias).Methods("GET")   // JSON com todas as categorias
	router.HandleFunc("/denuncias/{uri}", pegarUmaCategoria).Methods("GET") // devolve apenas uma categoria
	router.HandleFunc("/denuncias/", gravarNovaDenuncia).Methods("POST")    // adiciona nova denuncia

	log.Fatal(http.ListenAndServe(":8080", router)) // Server na porta 8080 [ localhost:8080 ]
}
