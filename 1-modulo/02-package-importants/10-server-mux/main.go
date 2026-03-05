package main

import (
	"encoding/json"
	"io"
	"net/http"
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
	mux := http.NewServeMux()
	mux.HandleFunc("/", BuscaCEPHandler)
	mux.Handle("blog", &blog{})
	http.ListenAndServe(":8080", mux)

	mux2 := http.NewServeMux()
	mux2.HandleFunc("/", BuscaCEPHandler)
	http.ListenAndServe(":8081", mux2)

}

type blog struct {
	title  string
	author string
}

func (b *blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, world"))
}

func BuscaCEPHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	data, err := BuscaCEP(cepParam)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func BuscaCEP(cep string) (*CepResponse, error) {
	resp, erro := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if erro != nil {
		return nil, erro
	}
	defer resp.Body.Close()

	body, erro := io.ReadAll(resp.Body)
	if erro != nil {
		return nil, erro
	}

	var cepResponse CepResponse
	erro = json.Unmarshal(body, &cepResponse)
	if erro != nil {
		return nil, erro
	}

	return &cepResponse, nil
}
