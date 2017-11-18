package paginas

import (
	"bufio"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"variaveis"
)

var PaginasEstaticas = populateStaticPages()

// função usada para atualizar os arquivos .html e .json
func AtualizarArquivosWeb() {
	PaginasEstaticas = populateStaticPages()
}

func populateStaticPages() *template.Template {
	resultado := template.New("templates")
	templatePaths := new([]string)

	templateFolder, _ := os.Open(variaveis.PastaPaginas)
	defer templateFolder.Close()
	templatePathsRaw, _ := templateFolder.Readdir(-1)
	for _, pathinfo := range templatePathsRaw {
		log.Println(pathinfo.Name())
		*templatePaths = append(*templatePaths, variaveis.PastaPaginas+"/"+pathinfo.Name())
	}

	basePath := variaveis.PastaTemas + variaveis.TemaDaPagina
	templateFolder, _ = os.Open(basePath)
	defer templateFolder.Close()
	templatePathsRaw, _ = templateFolder.Readdir(-1)
	for _, pathinfo := range templatePathsRaw {
		log.Println(pathinfo.Name())
		*templatePaths = append(*templatePaths, basePath+"/"+pathinfo.Name())
	}

	resultado.ParseFiles(*templatePaths...)
	return resultado
}

func ServeResource(w http.ResponseWriter, req *http.Request) {

	path := variaveis.Public + variaveis.TemaDaPagina + req.URL.Path
	var contentType string

	if strings.HasSuffix(path, ".css") {
		contentType = "text/css; charset=utf-8"
	} else if strings.HasSuffix(path, ".js") {
		contentType = "application/javascript; charset=utf-8"
	} else {
		contentType = "text/plain; charset=utf-8"
	}

	log.Println(path)
	f, err := os.Open(path)
	if err == nil {
		defer f.Close()
		w.Header().Add("Content-type", contentType)
		br := bufio.NewReader(f)
		br.WriteTo(w)
	} else {
		w.WriteHeader(404)
	}
}
