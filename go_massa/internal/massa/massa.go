package massa

const baseUrl = "http://v131d079.prevnet:8080/api"

type Massa struct {
	CpfMassa    string `json:"cpfMassa"`
	NitMassa    string `json:"nitMassa"`
	NomePfMassa string `json:"nomePfMassa"`
}

type Catalogo struct {
	Id                 int    `json:"id"`
	Descricao          string `json:"descricao"`
	Nome               string `json:"nome"`
	ChaveCatalogo      string `json:"chaveCatalogo"`
	NumeroEstoqueAtual int    `json:"numeroEstoqueAtual"`
	Visivel            string `json:"visivel"`
}
