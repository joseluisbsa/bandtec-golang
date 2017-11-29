// Todo codigo principal em Go precisa iniciar com este package
package main

// Importa todas as bibliotecas que serão utilizadas no código
import (
	"bd"
	"rest"
)
// funcao principal
func main() {
	bd.AtualizarTodasDenuncias()
	rest.IniciarServidor()
}
