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

func atualizaJSON() {
	log.Printf("atualiza arquivo JSON")
	// Busca na API todas as categorias e o total de denuncias de cada uma
	respostaFull, err := http.Get("http://localhost:8080/denuncias/")
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
	log.Println("Full: ", categFull)

	//for _, item := range categFull{
	//	log.Println(item.Nome, item.Total)
	//}

	///////////////////////////////////// busca por categoria e regiao

	respostaEach, err := http.Get("http://localhost:8080/denuncias/0")
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
	//for _, item := range categEach{
	//	log.Println(item.ID, item.Nome, item.Regiao, item.Total)
	//}

	//////////////////////////////////////// rotina para escrever nos arquivos
	// mudar o path !
	//path := "/home/joseph/github/bandtec-golang/poc-golang/versao_03/client-side/src/pages/"
	path := "C:/Users/aluno/bandtec-golang/poc-golang/versao_04/client-side/src/pages/"
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
			//categoriaFound = append(categoriaFound,item)
			//log.Println(item.ID, item.Nome, item.Regiao, item.Total)

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
	attHTML()
}

func attHTML() {
	// função usada para atualizar os arquivos .html e .json
	paginas.ThemeName = paginas.GetThemeName()
	paginas.StaticPages = paginas.PopulateStaticPages()
}
