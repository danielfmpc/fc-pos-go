package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID         uint    `gorm:"primarykey"`
	Name       string  `gorm:"size:255;not null"`
	Price      float64 `gorm:"not null"`
	CategoryID uint
	Category   Category
	gorm.Model
}

type Category struct {
	ID   uint   `gorm:"primarykey"`
	Name string `gorm:"size:255;not null"`
	gorm.Model
}

func NewProduct(name string, price float64, categoryID uint) *Product {
	return &Product{
		Name:       name,
		Price:      price,
		CategoryID: categoryID,
	}
}

func NewCategory(name string) *Category {
	return &Category{
		Name: name,
	}
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{})

	// category := NewCategory("Eletronicos")
	// db.Create(&category)
	db.Create(NewProduct("Notebook", 1500, 1))

	var products []Product
	db.Preload("Category").Find(&products)
	for _, p := range products {
		fmt.Println(p.Category.Name, "-", p.Name, "-", p.Price)
	}
}
