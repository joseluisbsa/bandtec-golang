package bd

import (
	"database/sql"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

// AbrirConexaoBanco inicia a conexão com o banco de dados
func AbrirConexaoBanco() {
	bancoDeDados, erro = sql.Open("mssql", stringDeConexao)
	verificarErro(erro, "ERRO AO CONECTAR COM BANCO DE DADOS", true)
}

func AtualizarTodasDenuncias() {
	AtualizarDenuncias()
	AtualizarDenunciasPorRegiao()
	AtualizarUltimoIDBanco()
}

// AtualizarDenuncias apenas busca no banco de dados e salva no array de structs 'Denuncias'
func AtualizarDenuncias() {

	AbrirConexaoBanco()
	// apaga dados existentes no array
	Denuncias = Denuncias[:0]

	query := `SELECT d.id_categoria, c.categoria, COUNT(d.id_categoria) 
			FROM tab_denuncia d JOIN tab_categoria c 
			ON d.id_categoria = c.id 
			GROUP BY d.id_categoria, c.categoria`

	consultarNoBanco(query, &Denuncias, false)
	// fecha conexão com banco
	defer bancoDeDados.Close()

	log.Printf("Denuncias atualizadas")
}

// AtualizarDenunciasPorRegiao apenas busca no banco de dados e salva no array de structs 'DenunciasPorRegiao'
func AtualizarDenunciasPorRegiao() {

	AbrirConexaoBanco()
	// apaga dados existentes no array
	DenunciasPorRegiao = DenunciasPorRegiao[:0]

	query := `SELECT d.id_categoria, c.categoria, COUNT(d.id_localidade), l.regiao
			FROM tab_denuncia d JOIN tab_categoria c 
			ON d.id_categoria = c.id 
			JOIN tab_localidade l 
			ON d.id_localidade = l.id 
			GROUP BY d.id_localidade, d.id_categoria, c.categoria, l.regiao`

	consultarNoBanco(query, &DenunciasPorRegiao, true)

	defer bancoDeDados.Close()

	log.Printf("Denuncias por regiao atualizadas")
}

// consultarNoBanco faz o select no banco de dados e salva em 'retornoSelectBanco'
func consultarNoBanco(query string, den *[]DadosDasDenuncias, porRegiao bool) {
	// SELECT principal
	retornoSelectBanco, erro := bancoDeDados.Query(query)

	if erro != nil {
		verificarErro(erro, "ERRO NO SELECT PRINCIPAL DAS DENUNCIAS", false)
	} else {
		for retornoSelectBanco.Next() {
			// 'addDenuncia' criado como uma struct vazia para receber cada linha retornada
			// do banco e depois ser adicionada no ponteiro do array das denuncias
			addDenuncia := DadosDasDenuncias{}
			// se for por regiao o select traz um campo a mais o addDenuncia.Regiao
			switch porRegiao {
			case false:
				erro := retornoSelectBanco.Scan(&addDenuncia.ID, &addDenuncia.Nome, &addDenuncia.Total)
				verificarErro(erro, "ERRO AO SALVAR AS DENUNCIAS RETORNADAS DO BANCO NA STRUCT", true)
			case true:
				erro := retornoSelectBanco.Scan(&addDenuncia.ID, &addDenuncia.Nome, &addDenuncia.Total, &addDenuncia.Regiao)
				verificarErro(erro, "ERRO AO SALVAR AS DENUNCIAS POR REGIAO RETORNADAS DO BANCO NA STRUCT", true)
			}
			// adiocina cada denuncia encontrada no ponteiro
			*den = append(*den, addDenuncia)
		}
	}
}

func AtualizarUltimoIDBanco() {

	AbrirConexaoBanco()
	ultimoIDBanco, erro := bancoDeDados.Query("SELECT MAX(id) FROM tab_denuncia")
	defer ultimoIDBanco.Close()
	defer bancoDeDados.Close()

	if erro != nil {
		verificarErro(erro, "ERRO NO SELECT MAX DE DENUNCIAS", false)
	} else {
		for ultimoIDBanco.Next() {
			// salve a quantidade de denuncias já existentes
			erro := ultimoIDBanco.Scan(&proximoIdParaGravarNoBanco)

			if erro != nil {
				verificarErro(erro, "ERRO AO SALVAR O ULTIMO ID NA VARIAVEL", false)
			} else {
				proximoIdParaGravarNoBanco++
			}
		}
	}
	log.Printf("Ultimo ID atualizado: %d", proximoIdParaGravarNoBanco)
}

func GravarNovaDenuncia(nova *NovaDenuncia) {

	AbrirConexaoBanco()
	gravar, erro := bancoDeDados.Query(`INSERT into tab_denuncia (id, id_categoria, id_localidade) 
											VALUES (?1, ?2, ?3)`, proximoIdParaGravarNoBanco, &nova.Categoria, &nova.Localidade)

	defer gravar.Close()       // fecha comando Query
	defer bancoDeDados.Close() // fecha conexão com o Banco

	if erro != nil {
		verificarErro(erro, "ERRO NO INSERT", false)
	} else {
		proximoIdParaGravarNoBanco++
	}
}
