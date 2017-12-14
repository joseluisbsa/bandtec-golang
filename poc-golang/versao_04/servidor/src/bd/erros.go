package bd

import "log"

func verificarErro(erro error, msg string, critico bool) {

	if erro != nil {
		switch critico {
		case false:
			// apenas exibe a mensagem e o erro na console e não para
			log.Println(msg + ": " + erro.Error())
		case true:
			// para totalmente a execução e exibe a mensagem e o erro
			log.Fatal(msg + ": " + erro.Error())
		}
	}
}
