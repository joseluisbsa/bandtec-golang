package bd

import (
	"database/sql"
	"fmt"
)

// array usado para enviar todas denuncias
var Denuncias []DadosDasDenuncias

// array usado para enviar as denuncias por regiao
var DenunciasPorRegiao []DadosDasDenuncias

// Usado para armazenar o ultimo 'id' do banco de dados
var proximoIdParaGravarNoBanco int
var erro error

// usado sempre que for preciso abrir conexão com o banco de dados
var stringDeConexao = fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s;port=%d", servidor, usuario, senha, banco, porta)

var bancoDeDados, _ = sql.Open("tipoBanco", "stringDeConexao")

type DadosDasDenuncias struct {
	ID     string `json:"id,omitempty"`
	Nome   string `json:"nome,omitempty"`
	Total  string `json:"total,omitempty"`
	Regiao string `json:"regiao,omitempty"`
}

type NovaDenuncia struct {
	Categoria  string `json:"categoria,omitempty"`
	Localidade string `json:"localidade,omitempty"`
}
