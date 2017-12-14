package servidor

import (
	"testing"
)

func TestRequisitarDados(t *testing.T) {

	for _, teste := range []struct {
		url string
	}{
		{urlTodasDenuncias},
		{urlTodasDenunciasPorRegiao},
	} {
		retornoTeste, erro := RequisitarDados(teste.url)
		if erro != nil {
			t.Fatalf("ERRO NA REQUISICAO DOS DADOS")
		}
		if retornoTeste == nil {
			t.Fatalf("VERIFIQUE SE AS URLs ESTAO CORRETAS")
		}

		for _, rt := range retornoTeste {
			t.Log(rt.ID, rt.Nome, rt.Total, rt.Regiao)
		}
		t.Log("-@-@-@-@-@-@-@-@-@-@-@-@-@-@-@-@-@-@-@-@-@-@-@-")
	}
}
