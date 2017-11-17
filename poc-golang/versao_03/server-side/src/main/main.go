// Todo codigo em Go precisa iniciar com este package
package main

// Importa todas as bibliotecas que serão utilizadas no código
import (
	"bd"
	"rest"
)

func main() {
	bd.AtualizarTodasDenuncias()
	rest.IniciarServidor()
}