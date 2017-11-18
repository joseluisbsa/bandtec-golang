package servidor

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"paginas"
	"strings"
	"variaveis"
)

func atualizarJSON() {
	log.Printf("atualiza arquivo JSON")
	// Busca na API todas as categorias e o total de denuncias de cada uma
	respostaFull, err := http.Get(variaveis.TodasDenuncias)
	if err != nil {
		log.Println(err)
	}

	// Copia apenas o corpo onde esta os dados solicitados
	dadosRespostaFull, err := ioutil.ReadAll(respostaFull.Body)
	if err != nil {
		log.Println(err)
	}
	// cria a Struct que ira salvar os dados recebidos da API
	var categFull []variaveis.CategoriaFull
	// coverte de json para struct
	json.Unmarshal(dadosRespostaFull, &categFull)

	///////////////////////////////////// busca por categoria e regiao

	respostaEach, err := http.Get(variaveis.TodasDenunciasPorRegiao)
	if err != nil {
		log.Println(err)
	}

	dadosRespostaEach, err := ioutil.ReadAll(respostaEach.Body)
	if err != nil {
		log.Println(err)
	}

	var categEach []variaveis.CategoriaEach
	json.Unmarshal(dadosRespostaEach, &categEach)
	log.Println("Full: ", categEach)

	//////////////////////////////////////// rotina para escrever nos arquivos
	// mudar o path !
	//path := "/home/joseph/github/bandtec-golang/poc-golang/versao_03/client-side/src/pages/"
	path := variaveis.ArquivosJSON
	// arquivo default.json com o formato padrão do JSON que a pagina lê
	jsonOut, err := ioutil.ReadFile(path + "default.json")
	if err != nil {
		log.Println(err)
	}

	for _, item := range categFull {
		// Com o conteudo lido do arquivo 'default.json', sera substituido a 1ª palavra 'Categoria'
		// pelo Nome da categoria salva em 'categFull' atual do for range
		attCategoria := bytes.Replace(jsonOut, []byte("Categoria"), []byte(item.Nome), 1)
		// Sobrescreve/Cria arquivo 'geral.json.html' com o conteudo de 'default.json'
		// em 'path' esta o camminho onde o arquivo deve ser salvo
		if err = ioutil.WriteFile(path+"geral.json.html", attCategoria, 0666); err != nil {
			log.Println(err)
		}
		// Lê o novo conteudo do JSON, caso contrario iria sobrescrever
		jsonOut, err = ioutil.ReadFile(path + "geral.json.html")
		if err != nil {
			log.Println(err)
		}

		attTotal := bytes.Replace(jsonOut, []byte("00"), []byte(item.Total), 1)
		if err = ioutil.WriteFile(path+"geral.json.html", attTotal, 0666); err != nil {
			log.Println(err)
		}

		jsonOut, err = ioutil.ReadFile(path + "geral.json.html")
		if err != nil {
			log.Println(err)
		}
	}

	jsonOut, err = ioutil.ReadFile(path + "default.json")
	if err != nil {
		log.Println(err)
	}
	// Mesma rotina acima, porem agora separado as categorias por região
	for _, item := range categEach {
		// Para comparar se os nomes são iguais deixo os dois em CAIXA ALTO e comparo.
		if strings.ToUpper(item.Nome) == strings.ToUpper(variaveis.PageAlias) {

			attCategoria := bytes.Replace(jsonOut, []byte("Categoria"), []byte(item.Regiao), 1)
			if err = ioutil.WriteFile(path+"categoria.json.html", attCategoria, 0666); err != nil {
				log.Println(err)
			}

			jsonOut, err = ioutil.ReadFile(path + "categoria.json.html")
			if err != nil {
				log.Println(err)
			}

			attTotal := bytes.Replace(jsonOut, []byte("00"), []byte(item.Total), 1)
			if err = ioutil.WriteFile(path+"categoria.json.html", attTotal, 0666); err != nil {
				log.Println(err)
			}

			jsonOut, err = ioutil.ReadFile(path + "categoria.json.html")
			if err != nil {
				log.Println(err)
			}
		}
	}
	// Atualiza todos os arquivos .html e .json que serão usados nas paginas
	paginas.AtualizarArquivosWeb()
}
