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

func LerArquivoJSON(arquivo string) ([]byte, error) {
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

	return dadosArquivoJSON, erro
}

func alterarArquivosJSON(denuncias []DadosDasDenuncias, arquivo string, verificaPorRegiao bool) {
	// ignoro o erro colocando o _
	dadosArquivoJSON, _ := LerArquivoJSON("modelo.json")
	// escreveArquivo usado quando for por regiao, se o nome da pagina selecionada
	// for igual o Nome do objeto corrente no FOR RANGE
	var escreveArquivo = true

	for _, item := range denuncias {
		alteraDescricao := []byte(item.Nome)
		if verificaPorRegiao == true {
			alteraDescricao = []byte(item.Regiao)
			escreveArquivo = false
			// coloca as duas strings em caixa alta para comparar
			if strings.ToUpper(item.Nome) == strings.ToUpper(paginaSelecionada) {
				escreveArquivo = true
			}
		}

		if escreveArquivo == true {
			// dadosArquivoJSON esta com o conteudo do modelo.json
			// o Replace busca pela primeira para "Categoria" e
			// substitui pelo Nome salvo em alteraDescricao
			jsonAlterado := bytes.Replace(dadosArquivoJSON, []byte("Categoria"), alteraDescricao, 1)
			escreverArquivoJSON(arquivo, jsonAlterado)

			dadosArquivoJSON, _ = LerArquivoJSON(arquivo)
			
			jsonAlterado = bytes.Replace(dadosArquivoJSON, []byte("00"), []byte(item.Total), 1)
			escreverArquivoJSON(arquivo, jsonAlterado)

			dadosArquivoJSON, _ = LerArquivoJSON(arquivo)
		}
	}
}

func RequisitarDados(url string) ([]DadosDasDenuncias, error) {

	var denuncias []DadosDasDenuncias
	// faz um GET no servidor rest
	requisicao, erro := http.Get(url)
	if erro != nil {
		return nil, erro
	}
	// armazena JSON retornado do GET
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
	// atualiza as paginas HTML
	atualizarArquivosWeb()
}
