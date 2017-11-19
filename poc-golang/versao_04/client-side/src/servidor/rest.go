package servidor

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func requisitarDados(url string) []DadosDasDenuncias {
	requisicao, erro := http.Get(url)
	if erro != nil {
		log.Println(erro)
	}

	corpoDaRequisicao, erro := ioutil.ReadAll(requisicao.Body)
	if erro != nil {
		log.Println(erro)
	}

	var denuncias []DadosDasDenuncias
	// coverte de json para struct
	json.Unmarshal(corpoDaRequisicao, &denuncias)

	return denuncias
}

func lerArquivoJSON(arquivo string) []byte {
	var jsonOut []byte
	var erro error
	switch arquivo {
	case "default.json":
		jsonOut, erro = ioutil.ReadFile(localArquivosJSON + arquivo)
	case "geral.json.html":
		jsonOut, erro = ioutil.ReadFile(localArquivosJSON + arquivo)
	case "categoria.json.html":
		jsonOut, erro = ioutil.ReadFile(localArquivosJSON + arquivo)

	}
	if erro != nil {
		log.Println(erro)
	}

	return jsonOut
}

func escreverArquivoJSON(denuncias []DadosDasDenuncias, arquivo string, verificaPorRegiao bool) {

	jsonOut := lerArquivoJSON("default.json")
	var continua = true

	for _, item := range denuncias {

		alteracao := []byte(item.Nome)

		if verificaPorRegiao == true {
			alteracao = []byte(item.Regiao)

			continua = false
			if strings.ToUpper(item.Nome) == strings.ToUpper(paginaSelecionada) {
				continua = true
			}
		}

		if continua == true {

			attCategoria := bytes.Replace(jsonOut, []byte("Categoria"), alteracao, 1)

			if erro := ioutil.WriteFile(localArquivosJSON+arquivo, attCategoria, 0666); erro != nil {
				log.Println(erro)
			}

			jsonOut = lerArquivoJSON(arquivo)

			attTotal := bytes.Replace(jsonOut, []byte("00"), []byte(item.Total), 1)
			if erro := ioutil.WriteFile(localArquivosJSON+arquivo, attTotal, 0666); erro != nil {
				log.Println(erro)
			}

			jsonOut = lerArquivoJSON(arquivo)
		}
	}
}

func atualizarArquivoJSON() {
	log.Printf("atualiza arquivo JSON")

	denuncias := requisitarDados(urlTodasDenuncias)
	escreverArquivoJSON(denuncias, "geral.json.html", false)

	denunciasPorRegiao := requisitarDados(urlTodasDenunciasPorRegiao)
	escreverArquivoJSON(denunciasPorRegiao, "categoria.json.html", true)

	atualizarArquivosWeb()
}
