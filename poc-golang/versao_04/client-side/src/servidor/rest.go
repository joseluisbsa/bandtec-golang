package servidor

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"variaveis"
)

func requisitarDados(url string) []variaveis.DadosDasDenuncias {
	requisicao, erro := http.Get(url)
	if erro != nil {
		log.Println(erro)
	}

	corpoDaRequisicao, erro := ioutil.ReadAll(requisicao.Body)
	if erro != nil {
		log.Println(erro)
	}

	var denuncias []variaveis.DadosDasDenuncias
	// coverte de json para struct
	json.Unmarshal(corpoDaRequisicao, &denuncias)

	return denuncias
}

func atualizarArquivoJSON() {
	log.Printf("atualiza arquivo JSON")

	denuncias := requisitarDados(variaveis.URLTodasDenuncias)
	denunciasPorRegiao := requisitarDados(variaveis.URLTodasDenunciasPorRegiao)

	// log.Println("Full: ", denunciasPorRegiao)
	// log.Println("------------------")
	// log.Println("Full: ", denuncias)

	//////////////////////////////////////// rotina para escrever nos arquivos
	path := variaveis.ArquivosJSON
	// arquivo default.json com o formato padrão do JSON que a pagina lê
	jsonOut, erro := ioutil.ReadFile(path + "default.json")
	if erro != nil {
		log.Println(erro)
	}

	for _, item := range denuncias {
		// Com o conteudo lido do arquivo 'default.json', sera substituido a 1ª palavra 'Categoria'
		// pelo Nome da categoria salva em 'denuncias' atual do for range
		attCategoria := bytes.Replace(jsonOut, []byte("Categoria"), []byte(item.Nome), 1)
		// Sobrescreve/Cria arquivo 'geral.json.html' com o conteudo de 'default.json'
		// em 'path' esta o camminho onde o arquivo deve ser salvo
		if erro = ioutil.WriteFile(path+"geral.json.html", attCategoria, 0666); erro != nil {
			log.Println(erro)
		}
		// Lê o novo conteudo do JSON, caso contrario iria sobrescrever
		jsonOut, erro = ioutil.ReadFile(path + "geral.json.html")
		if erro != nil {
			log.Println(erro)
		}

		attTotal := bytes.Replace(jsonOut, []byte("00"), []byte(item.Total), 1)
		if erro = ioutil.WriteFile(path+"geral.json.html", attTotal, 0666); erro != nil {
			log.Println(erro)
		}

		jsonOut, erro = ioutil.ReadFile(path + "geral.json.html")
		if erro != nil {
			log.Println(erro)
		}
	}

	jsonOut, erro = ioutil.ReadFile(path + "default.json")
	if erro != nil {
		log.Println(erro)
	}
	// Mesma rotina acima, porem agora separado as categorias por região
	for _, item := range denunciasPorRegiao {
		// Para comparar se os nomes são iguais deixo os dois em CAIXA ALTO e comparo.
		if strings.ToUpper(item.Nome) == strings.ToUpper(variaveis.PaginaSelecionada) {

			attCategoria := bytes.Replace(jsonOut, []byte("Categoria"), []byte(item.Regiao), 1)
			if erro = ioutil.WriteFile(path+"categoria.json.html", attCategoria, 0666); erro != nil {
				log.Println(erro)
			}

			jsonOut, erro = ioutil.ReadFile(path + "categoria.json.html")
			if erro != nil {
				log.Println(erro)
			}

			attTotal := bytes.Replace(jsonOut, []byte("00"), []byte(item.Total), 1)
			if erro = ioutil.WriteFile(path+"categoria.json.html", attTotal, 0666); erro != nil {
				log.Println(erro)
			}

			jsonOut, erro = ioutil.ReadFile(path + "categoria.json.html")
			if erro != nil {
				log.Println(erro)
			}
		}
	}
	// Atualiza todos os arquivos .html e .json que serão usados nas paginas
	atualizarArquivosWeb()
}
