package servidor

import "log"

func verificarErro(erro error) {
	if erro != nil {
		log.Println(erro)
	}
}
