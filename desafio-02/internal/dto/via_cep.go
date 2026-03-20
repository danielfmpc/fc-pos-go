package dto

type ViaCep struct {
	Cep        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Cidade     string `json:"localidade"`
	Estado     string `json:"estado"`
}

func (v *ViaCep) ToCep() Cep {
	return Cep{
		Cep:        v.Cep,
		Logradouro: v.Logradouro,
		Bairro:     v.Bairro,
		Cidade:     v.Cidade,
		Estado:     v.Estado,
		Source:     "via_cep",
	}
}
