package bd

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
