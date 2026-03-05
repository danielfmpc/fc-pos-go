package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type CepResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, arg := range os.Args[1:] {
		req, err := http.Get("https://viacep.com.br/ws/" + arg + "/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err)
		}
		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler resposta: %v\n", err)
		}

		var viaCep CepResponse

		err = json.Unmarshal(res, &viaCep)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao decodificar JSON: %v\n", err)
		}

		fmt.Println(viaCep)

		file, err := os.Create("Cidade-" + viaCep.Localidade + "-" + viaCep.Uf + ".json")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao criar arquivo: %v\n", err)
		}
		defer file.Close()

		_, err = file.Write([]byte(res))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao escrever no arquivo: %v\n", err)
		}
	}
}
