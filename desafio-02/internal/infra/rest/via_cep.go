package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/danielfmpc/pos-go-expert-desafio02/internal/dto"
)

type ViaCepClient struct {
}

func (v *ViaCepClient) GetCep(ctx context.Context, cep string) (*dto.Cep, error) {
	viaCep, err := callViaCep(ctx, cep)
	if err != nil {
		return nil, err
	}

	result := viaCep.ToCep()
	return &result, nil
}

func callViaCep(ctx context.Context, cep string) (*dto.ViaCep, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var viaCep dto.ViaCep
	err = json.NewDecoder(res.Body).Decode(&viaCep)
	if err != nil {
		return nil, err
	}

	return &viaCep, nil
}
