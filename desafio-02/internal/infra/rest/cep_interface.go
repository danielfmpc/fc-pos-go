package rest

import (
	"context"

	"github.com/danielfmpc/pos-go-expert-desafio02/internal/dto"
)

type CepInterface interface {
	GetCep(ctx context.Context, cep string) (*dto.Cep, error)
}
