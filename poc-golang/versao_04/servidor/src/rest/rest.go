package rest

import (
	"bd"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// devolve todas denuncias via (localhost:8080/denuncias/)
func pegarTodasDenuncias(w http.ResponseWriter, req *http.Request) {
	log.Printf("GET todas denuncias")
	json.NewEncoder(w).Encode(bd.Denuncias)
}

// devolve todas denuncias separado por regiao (localhost:8080/denuncias/0)
func pegarDenunciasPorRegiao(w http.ResponseWriter, req *http.Request) {
	log.Printf("GET denuncias por regiao")
	parametros := mux.Vars(req)
	var categoriaEncontrada []bd.DadosDasDenuncias
	for _, item := range bd.DenunciasPorRegiao {
		if strings.ToLower(item.Nome) == strings.ToLower(parametros["uri"]) {
			categoriaEncontrada = append(categoriaEncontrada, item)
		}
	}
	json.NewEncoder(w).Encode(bd.DenunciasPorRegiao)
}

// Adicona mais uma denuncia
func gravarNovaDenuncia(w http.ResponseWriter, req *http.Request) {
	// modelo que deve enviado
	// {"categoria":"4","localidade":"2"}
	log.Printf("POST nova Denuncia")
	var NovaD bd.NovaDenuncia

	// grava em 'novaD' os dados recebidos
	erro := json.NewDecoder(req.Body).Decode(&NovaD)
	verificarErro(erro, "ERRO AO SALVAR OS DADOS RECEBIDOS EM 'NOVAD'", false)

	// imprime no terminal os valores recebidos
	log.Println(NovaD)
	bd.GravarNovaDenuncia(&NovaD)
	bd.AtualizarTodasDenuncias()
}
