package bd

import (
	"database/sql"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

// array usado para enviar o total de cada categoria
var Denuncias []DadosDasDenuncias

// array usado para enviar o total de denuncias por regiao
var DenunciasPorCategoria []DadosDasDenuncias

// Usado para armazenar o ultimo 'id' do banco de dados
var proximoIdParaGravarNoBanco int

func GravarNovaDenuncia(nova *NovaDenuncia) {

	AbrirConexaoBanco()
	gravar, erro := bancoDeDados.Query(`INSERT into tab_denuncia (id, id_categoria, id_localidade) 
										VALUES (?1, ?2, ?3)`, proximoIdParaGravarNoBanco, &nova.Categoria, &nova.Localidade)

	defer gravar.Close()       // fecha comando Query
	defer bancoDeDados.Close() // fecha conexão com o Banco

	if erro != nil {
		log.Println("erro no INSERT:", erro.Error())
	} else {
		proximoIdParaGravarNoBanco++
	}
}

// AbrirConexaoBanco inicia a conexão com o banco de dados
func AbrirConexaoBanco() {
	bancoDeDados, erro = sql.Open("mssql", stringDeConexao)
	verificarErro(erro, "Erro ao conectar ao banco de dados", true)
}

func AtualizarTodasDenuncias() {
	AtualizarDenuncias()
	AtualizarDenunciasPorCategoria()
	AtualizarUltimoIDBanco()
}

func AtualizarDenuncias() {

	AbrirConexaoBanco()

	Denuncias = Denuncias[:0]

	query := `SELECT d.id_categoria, c.categoria, COUNT(d.id_categoria) 
			FROM tab_denuncia d JOIN tab_categoria c 
			ON d.id_categoria = c.id 
			GROUP BY d.id_categoria, c.categoria`

	consultarNoBanco(query, &Denuncias, false)

	defer bancoDeDados.Close()

	log.Printf("Denuncias atualizadas")
}

func AtualizarDenunciasPorCategoria() {

	AbrirConexaoBanco()

	DenunciasPorCategoria = DenunciasPorCategoria[:0]

	query := `SELECT d.id_categoria, c.categoria, COUNT(d.id_localidade), l.regiao
			FROM tab_denuncia d JOIN tab_categoria c 
			ON d.id_categoria = c.id 
			JOIN tab_localidade l 
			ON d.id_localidade = l.id 
			GROUP BY d.id_localidade, d.id_categoria, c.categoria, l.regiao`

	consultarNoBanco(query, &DenunciasPorCategoria, true)

	defer bancoDeDados.Close()

	log.Printf("Denuncias por categorias atualizadas")
}

func consultarNoBanco(query string, den *[]DadosDasDenuncias, porRegiao bool) {

	retornoSelectBanco, erro := bancoDeDados.Query(query)

	if erro != nil {
		verificarErro(erro, "Erro no select das denuncias", false)
	} else {
		for retornoSelectBanco.Next() {
			addCategoria := DadosDasDenuncias{}
			if porRegiao == false {
				if erro := retornoSelectBanco.Scan(&addCategoria.ID, &addCategoria.Nome, &addCategoria.Total); erro != nil {
					log.Println("erro ao salvar as categoriasFull retornados do Banco:", erro.Error())
				}
			} else {
				if erro := retornoSelectBanco.Scan(&addCategoria.ID, &addCategoria.Nome, &addCategoria.Total, &addCategoria.Regiao); erro != nil {
					log.Println("erro ao salvar as categoriasEach retornados do Banco:", erro.Error())
				}
			}
			*den = append(*den, addCategoria)
		}
	}
}

func AtualizarUltimoIDBanco() {

	AbrirConexaoBanco()
	ultimoIDBanco, erro := bancoDeDados.Query("select MAX(id) from tab_denuncia")
	defer ultimoIDBanco.Close()
	defer bancoDeDados.Close()

	if erro != nil {
		verificarErro(erro, "Erro no select MAX de denuncias", false)
	} else {
		for ultimoIDBanco.Next() {

			if erro := ultimoIDBanco.Scan(&proximoIdParaGravarNoBanco); erro != nil {
				log.Println("erro ao salvar categoriasCount retornados do Banco:", erro.Error())
			} else {
				proximoIdParaGravarNoBanco++
			}
		}
	}
	log.Printf("Ultimo ID atualizado: %d", proximoIdParaGravarNoBanco)
}
