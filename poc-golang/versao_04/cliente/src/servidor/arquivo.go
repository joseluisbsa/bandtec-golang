package servidor

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func escreverArquivoJSON(arquivo string, jsonAlterado []byte) {
	if erro := ioutil.WriteFile(localArquivosHTMLeJSON+"/"+arquivo, jsonAlterado, 0666); erro != nil {
		log.Println(erro)
	}
}

func lerArquivoJSON(arquivo string) []byte {
	var dadosArquivoJSON []byte
	var erro error

	switch arquivo {
	case "modelo.json":
		dadosArquivoJSON, erro = ioutil.ReadFile(localArquivosHTMLeJSON + "/" + arquivo)
	case "geral.json.html":
		dadosArquivoJSON, erro = ioutil.ReadFile(localArquivosHTMLeJSON + "/" + arquivo)
	case "categoria.json.html":
		dadosArquivoJSON, erro = ioutil.ReadFile(localArquivosHTMLeJSON + "/" + arquivo)
	}
	verificarErro(erro, "ERRO AO LER ARQUIVO", false)

	return dadosArquivoJSON
}

func alterarArquivosJSON(denuncias []DadosDasDenuncias, arquivo string, verificaPorRegiao bool) {

	dadosArquivoJSON := lerArquivoJSON("modelo.json")
	var escreveArquivo = true

	for _, item := range denuncias {

		alteraDescricao := []byte(item.Nome)

		if verificaPorRegiao == true {
			alteraDescricao = []byte(item.Regiao)
			escreveArquivo = false

			if strings.ToUpper(item.Nome) == strings.ToUpper(paginaSelecionada) {
				escreveArquivo = true
			}
		}

		if escreveArquivo == true {
			jsonAlterado := bytes.Replace(dadosArquivoJSON, []byte("Categoria"), alteraDescricao, 1)
			escreverArquivoJSON(arquivo, jsonAlterado)

			dadosArquivoJSON = lerArquivoJSON(arquivo)

			jsonAlterado = bytes.Replace(dadosArquivoJSON, []byte("00"), []byte(item.Total), 1)
			escreverArquivoJSON(arquivo, jsonAlterado)

			dadosArquivoJSON = lerArquivoJSON(arquivo)
		}
	}
}

func requisitarDados(url string) []DadosDasDenuncias {

	var denuncias []DadosDasDenuncias

	requisicao, erro := http.Get(url)
	verificarErro(erro, "ERRO NA REQUISICAO DOS DADOS - GET", false)

	corpoDaRequisicao, erro := ioutil.ReadAll(requisicao.Body)
	verificarErro(erro, "ERRO AO GRAVAR OS DADOS RECEBIDOS DA REQUISIÇÃO", false)
	// coverte de json para struct
	json.Unmarshal(corpoDaRequisicao, &denuncias)

	return denuncias
}

func atualizarArquivosJSON() {
	log.Printf("atualiza arquivo JSON")

	denuncias := requisitarDados(urlTodasDenuncias)
	alterarArquivosJSON(denuncias, "geral.json.html", false)

	denunciasPorRegiao := requisitarDados(urlTodasDenunciasPorRegiao)
	alterarArquivosJSON(denunciasPorRegiao, "categoria.json.html", true)

	atualizarArquivosWeb()
}
