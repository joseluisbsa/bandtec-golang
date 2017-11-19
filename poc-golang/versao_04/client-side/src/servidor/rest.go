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
		//log.Println(jsonOut)
		//log.Println("------------- ler arquivos")
	case "categoria.json.html":
		jsonOut, erro = ioutil.ReadFile(localArquivosJSON + arquivo)

	}
	if erro != nil {
		log.Println(erro)
	}
	//log.Println(jsonOut)
	//log.Println("------------- ler arquivos")

	return jsonOut
}

func escreverArquivoJSON(denuncias []DadosDasDenuncias, arquivo string) {

	jsonOut := lerArquivoJSON("default.json")
	//log.Println(jsonOut)
	//log.Println("################# volta do ler arquivos")

	for _, item := range denuncias {
		alteracao := []byte(item.Nome)
		attCategoria := bytes.Replace(jsonOut, []byte("Categoria"), alteracao, 1)

		if erro := ioutil.WriteFile(localArquivosJSON+arquivo, attCategoria, 0666); erro != nil {
			log.Println(erro)
		}

		jsonOut = lerArquivoJSON(arquivo)
		log.Println(jsonOut)
		log.Println("################# volta do ler arquivos")

		attTotal := bytes.Replace(jsonOut, []byte("00"), []byte(item.Total), 1)
		if erro := ioutil.WriteFile(localArquivosJSON+arquivo, attTotal, 0666); erro != nil {
			log.Println(erro)
		}

		jsonOut = lerArquivoJSON(arquivo)
	}
}

func atualizarArquivoJSON() {
	log.Printf("atualiza arquivo JSON")

	denuncias := requisitarDados(urlTodasDenuncias)

	//////////////////////////////////////// rotina para escrever nos arquivos
	// arquivo default.json com o formato padrão do JSON que a pagina lê
	//jsonOut := lerArquivoJSON("default.json")
	//log.Println(jsonOut)
	//log.Println("################# volta do ler arquivos")

	escreverArquivoJSON(denuncias, "geral.json.html")
	//log.Println(escreverArquivoJSON)
	//log.Println("################# volta")

	// for _, item := range denuncias {
	// 	// Com o conteudo lido do arquivo 'default.json', sera substituido a 1ª palavra 'Categoria'
	// 	// pelo Nome da categoria salva em 'denuncias' atual do for range
	// 	attCategoria := bytes.Replace(jsonOut, []byte("Categoria"), []byte(item.Nome), 1)
	// 	// Sobrescreve/Cria arquivo 'geral.json.html' com o conteudo de 'default.json'
	// 	// em 'localArquivosJSON' esta o camminho onde o arquivo deve ser salvo
	// 	if erro = ioutil.WriteFile(localArquivosJSON+"geral.json.html", attCategoria, 0666); erro != nil {
	// 		log.Println(erro)
	// 	}
	// 	// Lê o novo conteudo do JSON, caso contrario iria sobrescrever
	// 	jsonOut, erro = ioutil.ReadFile(localArquivosJSON + "geral.json.html")
	// 	if erro != nil {
	// 		log.Println(erro)
	// 	}

	// 	attTotal := bytes.Replace(jsonOut, []byte("00"), []byte(item.Total), 1)
	// 	if erro = ioutil.WriteFile(localArquivosJSON+"geral.json.html", attTotal, 0666); erro != nil {
	// 		log.Println(erro)
	// 	}

	// 	jsonOut, erro = ioutil.ReadFile(localArquivosJSON + "geral.json.html")
	// 	if erro != nil {
	// 		log.Println(erro)
	// 	}
	// }

	denunciasPorRegiao := requisitarDados(urlTodasDenunciasPorRegiao)

	jsonOut, erro := ioutil.ReadFile(localArquivosJSON + "default.json")
	if erro != nil {
		log.Println(erro)
	}
	// Mesma rotina acima, porem agora separado as categorias por região

	for _, item := range denunciasPorRegiao {
		// Para comparar se os nomes são iguais deixo os dois em CAIXA ALTO e comparo.
		if strings.ToUpper(item.Nome) == strings.ToUpper(paginaSelecionada) {

			attCategoria := bytes.Replace(jsonOut, []byte("Categoria"), []byte(item.Regiao), 1)
			if erro := ioutil.WriteFile(localArquivosJSON+"categoria.json.html", attCategoria, 0666); erro != nil {
				log.Println(erro)
			}

			jsonOut, erro = ioutil.ReadFile(localArquivosJSON + "categoria.json.html")
			if erro != nil {
				log.Println(erro)
			}

			attTotal := bytes.Replace(jsonOut, []byte("00"), []byte(item.Total), 1)
			if erro = ioutil.WriteFile(localArquivosJSON+"categoria.json.html", attTotal, 0666); erro != nil {
				log.Println(erro)
			}

			jsonOut, erro = ioutil.ReadFile(localArquivosJSON + "categoria.json.html")
			if erro != nil {
				log.Println(erro)
			}
		}
	}
	// Atualiza todos os arquivos .html e .json que serão usados nas paginas
	atualizarArquivosWeb()
}
