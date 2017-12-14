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
	erro := ioutil.WriteFile(localArquivosHTMLeJSON+"/"+arquivo, jsonAlterado, 0666)
	verificarErro(erro, "ERRO AO ESCREVER NO ARQUIVO JSON", false)
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

func RequisitarDados(url string) ([]DadosDasDenuncias, error) {

	var denuncias []DadosDasDenuncias

	requisicao, erro := http.Get(url)
	if erro != nil {
		return nil, erro
	}

	corpoDaRequisicao, erro := ioutil.ReadAll(requisicao.Body)
	if erro != nil {
		return nil, erro
	}
	// coverte de json para struct
	json.Unmarshal(corpoDaRequisicao, &denuncias)

	return denuncias, erro
}

func atualizarArquivosJSON() {
	log.Printf("atualiza arquivo JSON")

	denuncias, erro := RequisitarDados(urlTodasDenuncias)
	verificarErro(erro, "ERRO AO REQUISITAR DADOS VIA GET TODAS DENUNCIAS", false)
	alterarArquivosJSON(denuncias, "geral.json.html", false)

	denunciasPorRegiao, erro := RequisitarDados(urlTodasDenunciasPorRegiao)
	verificarErro(erro, "ERRO AO REQUISITAR DADOS VIA GET DENUNCIAS POR REGIAO", false)
	alterarArquivosJSON(denunciasPorRegiao, "categoria.json.html", true)

	atualizarArquivosWeb()
}
