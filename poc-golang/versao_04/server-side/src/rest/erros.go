package rest

import "log"

func verificarErro(erro error, msg string) {
	if erro != nil {
		log.Println(msg + ": " + erro.Error())
	}
}
