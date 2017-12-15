package servidor

import (
	"testing"
)

func TestLerArquivoJSON(t *testing.T) {
	for _, teste := range []struct {
		Arquivo string
	}{
		{"modelo.json"},
		{"geral.json.html"},
		{"categoria.json.html"},
	} {
		retornoTeste, erro := LerArquivoJSON(teste.Arquivo)
		if erro != nil {
			t.Fatalf("ERRO AO LER O ARQUIVO" + teste.Arquivo)
		}
		// converte BYTE para STRING
		t.Log(string(retornoTeste[:]))
		t.Log("-@-@-@-@-@-@-@-@-@-@-@-@-@-@-@-@-@-@-@-@-@-@-@-")
	}
}
func TestRequisitarDados(t *testing.T) {

	for _, teste := range []struct {
		Url string
	}{
		// valores que serão usados no teste
		{urlTodasDenuncias},
		{urlTodasDenunciasPorRegiao},
	} {
		retornoTeste, erro := RequisitarDados(teste.Url)
		if erro != nil {
			t.Fatalf("ERRO AO REQUISITAR DADOS DO SERVIDOR REST: " + teste.Url)
		}
		if retornoTeste == nil {
			t.Fatalf("SEM RETORNO PARA O GET: " + teste.Url)
		}
		// imprime na console o retorno
		for _, rt := range retornoTeste {
			t.Log(rt.ID, rt.Nome, rt.Total, rt.Regiao)
		}
		t.Log("-@-@-@-@-@-@-@-@-@-@-@-@-@-@-@-@-@-@-@-@-@-@-@-")
	}
}
