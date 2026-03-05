package main

import (
	"fmt"

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

	// var product Product
	// db.First(&product, 2)
	// fmt.Println(product)
	// db.First(&product, "name = ?", "Notebook")
	// fmt.Println(product)

	// var products []Product
	// db.Find(&products)
	// Paginação
	// db.Limit(2).Offset(2).Find(&products)
	// for _, p := range products {
	// 	fmt.Println(p)
	// }

	var p []Product
	// db.Where("price > ?", 10).Find(&p)
	// for _, p := range p {
	// 	fmt.Println(p)
	// }

	db.Where("name like ?", "%book%").Find(&p)
	for _, p := range p {
		fmt.Println(p)
	}
}
