package rest

import (
	"bd"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// Adicona mais uma denuncia
func gravarNovaDenuncia(w http.ResponseWriter, req *http.Request) {
	// modelo que deve enviado
	// {"categoria":"4","localidade":"2"}
	log.Printf("Post mais uma Nova Denuncia")
	var NovaD bd.NovaDenuncia

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
	var categoriaEncontrada []bd.DadosDasDenuncias
	for _, item := range bd.DenunciasPorCategoria {
		if strings.ToLower(item.Nome) == strings.ToLower(parametros["uri"]) {
			categoriaEncontrada = append(categoriaEncontrada, item)
		}
	}
	// envia todas por regiao, será filtrado no client-side
	json.NewEncoder(w).Encode(bd.DenunciasPorCategoria)
	// se quiser apenas uma categoria mesmo, comente a linha acima e descomente a de baixo
	//json.NewEncoder(w).Encode(categoriaEncontrada)
}

// envia os dados das categorias via GET
func pegarTodasCategorias(w http.ResponseWriter, req *http.Request) {
	log.Printf("Get categorias")
	json.NewEncoder(w).Encode(bd.Denuncias)
}
