package variaveis

var PageAlias string

type DefaultContext struct {
	Title string
}

type CategoriaFull struct {
	ID    string `json:"id,omitempty"`
	Nome  string `json:"nome,omitempty"`
	Total string `json:"total,omitempty"`
}

type CategoriaEach struct {
	ID     string `json:"id,omitempty"`
	Nome   string `json:"nome,omitempty"`
	Regiao string `json:"regiao,omitempty"`
	Total  string `json:"total,omitempty"`
}
