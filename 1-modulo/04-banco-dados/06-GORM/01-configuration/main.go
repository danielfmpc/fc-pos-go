package main

import (
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    string  `gorm:"primaryKey"`
	Name  string  `gorm:"size:255;not null"`
	Price float64 `gorm:"not null"`
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{})

	db.Create(NewProduct("Notebook", 1000.97))

	products := []*Product{
		NewProduct("Mouse", 150.00),
		NewProduct("Teclado", 200.00),
		NewProduct("Monitor", 900.00),
	}

	db.Create(&products)
}
