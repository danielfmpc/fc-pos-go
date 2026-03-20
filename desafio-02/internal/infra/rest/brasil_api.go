package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/danielfmpc/pos-go-expert-desafio02/internal/dto"
)

type BrasilApiClient struct {
}

func (b *BrasilApiClient) GetCep(ctx context.Context, cep string) (*dto.Cep, error) {
	brasilAPICep, err := callBrasilAPI(ctx, cep)
	if err != nil {
		return nil, err
	}

	result := brasilAPICep.ToCep()
	return &result, nil
}

func callBrasilAPI(ctx context.Context, cep string) (*dto.BrasilAPICep, error) {
	url := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var brasilAPICep dto.BrasilAPICep
	err = json.NewDecoder(res.Body).Decode(&brasilAPICep)
	if err != nil {
		return nil, err
	}

	return &brasilAPICep, nil
}
