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
	insert, erro := BancoDeDados.Query(`INSERT into tab_denuncia (id, id_categoria, id_localidade) 
										VALUES (?1, ?2, ?3)`, proximoIdParaGravarNoBanco, &nova.Categoria, &nova.Localidade)

	defer insert.Close()       // fecha comando Query
	defer BancoDeDados.Close() // fecha conexão com o Banco
	if erro != nil {
		log.Println("erro no INSERT:", erro.Error())
	} else {
		proximoIdParaGravarNoBanco++
	}
}

func AbrirConexaoBanco() {
	BancoDeDados, ErroBD = sql.Open("mssql", stringDeConexao)
	if ErroBD != nil {
		log.Println("erro ao conectar ao Banco: ", ErroBD.Error())
	}
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

	consultarNoBanco(query, &Denuncias, 1)

	defer BancoDeDados.Close()

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

	consultarNoBanco(query, &DenunciasPorCategoria, 2)

	defer BancoDeDados.Close()

	log.Printf("Denuncias por categorias atualizadas")
}

func consultarNoBanco(query string, den *[]DadosDasDenuncias, opcao int) {

	retornoSelectBanco, erro := BancoDeDados.Query(query)
	if erro != nil {
		log.Println("erro no SELECT das Categorias:", erro.Error())
	}

	for retornoSelectBanco.Next() {
		addCategoria := DadosDasDenuncias{}
		if opcao == 1 {
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

func AtualizarUltimoIDBanco() {

	AbrirConexaoBanco()

	ultimoIDBanco, erro := BancoDeDados.Query("select MAX(id) from tab_denuncia")
	if erro != nil {
		log.Println("erro no SELECT count categoria:", erro.Error())
	}

	defer ultimoIDBanco.Close()
	defer BancoDeDados.Close()

	for ultimoIDBanco.Next() {

		if erro := ultimoIDBanco.Scan(&proximoIdParaGravarNoBanco); erro != nil {
			log.Println("erro ao salvar categoriasCount retornados do Banco:", erro.Error())
		} else {
			proximoIdParaGravarNoBanco++
		}
	}
	log.Printf("Ultimo ID atualizado: %d", proximoIdParaGravarNoBanco)
}
