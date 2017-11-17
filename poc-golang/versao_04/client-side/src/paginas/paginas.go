package paginas

import (
	"html/template"
	"log"
	"os"
)

var ThemeName = GetThemeName()
var StaticPages = PopulateStaticPages()

func GetThemeName() string {
	return "bs4"
}

func PopulateStaticPages() *template.Template {
	result := template.New("templates")
	templatePaths := new([]string)

	//basePath := "/home/joseph/github/bandtec-golang/poc-golang/versao_03/client-side/src/pages"
	basePath := "C:/Users/aluno/bandtec-golang/poc-golang/versao_04/client-side/src/pages"
	templateFolder, _ := os.Open(basePath)
	defer templateFolder.Close()
	templatePathsRaw, _ := templateFolder.Readdir(-1)
	for _, pathinfo := range templatePathsRaw {
		log.Println(pathinfo.Name())
		*templatePaths = append(*templatePaths, basePath+"/"+pathinfo.Name())
	}

	//basePath = "/home/joseph/github/bandtec-golang/poc-golang/versao_03/client-side/src/themes/" + themeName
	basePath = "C:/Users/aluno/bandtec-golang/poc-golang/versao_04/client-side/src/themes/" + ThemeName
	templateFolder, _ = os.Open(basePath)
	defer templateFolder.Close()
	templatePathsRaw, _ = templateFolder.Readdir(-1)
	for _, pathinfo := range templatePathsRaw {
		log.Println(pathinfo.Name())
		*templatePaths = append(*templatePaths, basePath+"/"+pathinfo.Name())
	}

	result.ParseFiles(*templatePaths...)
	return result
}
