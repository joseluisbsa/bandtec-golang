package servidor

import (
	"bufio"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

var paginasEstaticas = populateStaticPages()

// função usada para atualizar os arquivos .html e .json
func atualizarArquivosWeb() {
	paginasEstaticas = populateStaticPages()
}

func populateStaticPages() *template.Template {
	resultado := template.New("templates")
	templatePaths := new([]string)

	templateFolder, _ := os.Open(pastaPaginas)
	defer templateFolder.Close()
	templatePathsRaw, _ := templateFolder.Readdir(-1)
	for _, pathinfo := range templatePathsRaw {
		log.Println(pathinfo.Name())
		*templatePaths = append(*templatePaths, pastaPaginas+"/"+pathinfo.Name())
	}

	basePath := pastaTemas + temaDaPagina
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

func serveResource(w http.ResponseWriter, req *http.Request) {

	path := public + temaDaPagina + req.URL.Path
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
