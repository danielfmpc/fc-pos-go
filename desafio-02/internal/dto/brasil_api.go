package dto

type BrasilAPICep struct {
	Cep        string `json:"cep"`
	Logradouro string `json:"street"`
	Bairro     string `json:"neighborhood"`
	Cidade     string `json:"city"`
	Estado     string `json:"state"`
}

func (b *BrasilAPICep) ToCep() Cep {
	return Cep{
		Cep:        b.Cep,
		Logradouro: b.Logradouro,
		Bairro:     b.Bairro,
		Cidade:     b.Cidade,
		Estado:     b.Estado,
		Source:     "brasil_api",
	}
}
