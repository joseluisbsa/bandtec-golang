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

func atualizarArquivoJSON() {
	log.Printf("atualiza arquivo JSON")
	// Busca na API todas as categorias e o total de denuncias de cada uma
	respostaFull, erro := http.Get(variaveis.TodasDenuncias)
	if erro != nil {
		log.Println(erro)
	}

	// Copia apenas o corpo onde esta os dados solicitados
	dadosRespostaFull, erro := ioutil.ReadAll(respostaFull.Body)
	if erro != nil {
		log.Println(erro)
	}
	// cria a Struct que ira salvar os dados recebidos da API
	var categFull []variaveis.CategoriaFull
	// coverte de json para struct
	json.Unmarshal(dadosRespostaFull, &categFull)

	///////////////////////////////////// busca por categoria e regiao

	respostaEach, erro := http.Get(variaveis.TodasDenunciasPorRegiao)
	if erro != nil {
		log.Println(erro)
	}

	dadosRespostaEach, erro := ioutil.ReadAll(respostaEach.Body)
	if erro != nil {
		log.Println(erro)
	}

	var categEach []variaveis.CategoriaEach
	json.Unmarshal(dadosRespostaEach, &categEach)
	log.Println("Full: ", categEach)

	//////////////////////////////////////// rotina para escrever nos arquivos
	path := variaveis.ArquivosJSON
	// arquivo default.json com o formato padrão do JSON que a pagina lê
	jsonOut, erro := ioutil.ReadFile(path + "default.json")
	if erro != nil {
		log.Println(erro)
	}

	for _, item := range categFull {
		// Com o conteudo lido do arquivo 'default.json', sera substituido a 1ª palavra 'Categoria'
		// pelo Nome da categoria salva em 'categFull' atual do for range
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
	for _, item := range categEach {
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
