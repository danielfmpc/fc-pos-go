//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/danielfmpc/pos-go-di/product"
	"github.com/google/wire"
)

var setRepoDependency = wire.NewSet(
	product.NewProductRepository,
	wire.Bind(new(product.ProductRepositoryInterface), new(*product.ProductRepository)),
)

func NewUseCase(db *sql.DB) *product.ProductUseCase {
	wire.Build(
		setRepoDependency,
		product.NewProductUseCase,
	)

	return &product.ProductUseCase{}
}
