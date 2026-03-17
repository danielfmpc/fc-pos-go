package database

import (
	"fmt"
	"math/rand"
	"pos-go-api/internal/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Product 1", 10)
	assert.NoError(t, err)

	productDB := NewProduct(db)
	assert.NoError(t, productDB.Create(product))

	productFound, err := productDB.FindByID(product.ID.String())
	assert.NoError(t, err)

	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}

func TestFindAllProducts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	productDb := NewProduct(db)
	for i := 1; i <= 24; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		assert.NoError(t, err)
		db.Create(product)
	}

	products, err := productDb.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)

	products, err = productDb.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 11", products[0].Name)
	assert.Equal(t, "Product 20", products[9].Name)

	products, err = productDb.FindAll(3, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 4)
	assert.Equal(t, "Product 21", products[0].Name)
	assert.Equal(t, "Product 24", products[3].Name)
}

func TestFindProductByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Product 1", 10.0)
	assert.NoError(t, err)
	db.Create(product)

	productDB := NewProduct(db)
	productFound, err := productDB.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}

func TestUpdateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Product 1", 10.0)
	assert.NoError(t, err)
	db.Create(product)

	productDB := NewProduct(db)
	product.Name = "Product 1 Updated"
	product.Price = 20.0
	assert.NoError(t, productDB.Update(product))

	productFound, err := productDB.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, "Product 1 Updated", productFound.Name)
	assert.Equal(t, 20.0, productFound.Price)
}

func TestDeleteProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Product 1", 10.0)
	assert.NoError(t, err)
	db.Create(product)

	productDB := NewProduct(db)
	assert.NoError(t, productDB.Delete(product.ID.String()))

	_, err = productDB.FindByID(product.ID.String())
	assert.Error(t, err)
}
