package servidor

import "log"

// verificarErro comum para todos os programas
func verificarErro(erro error) {
	if erro != nil {
		log.Println(erro)
	}
}
