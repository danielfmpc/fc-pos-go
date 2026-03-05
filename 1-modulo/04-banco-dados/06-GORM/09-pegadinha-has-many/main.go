package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID           uint    `gorm:"primarykey"`
	Name         string  `gorm:"size:255;not null"`
	Price        float64 `gorm:"not null"`
	CategoryID   uint
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type Category struct {
	ID       uint   `gorm:"primarykey"`
	Name     string `gorm:"size:255;not null"`
	Products []Product
	gorm.Model
}

type SerialNumber struct {
	ID        uint   `gorm:"primarykey"`
	Number    string `gorm:"size:255;not null"`
	ProductID uint
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

func NewSerialNumber(number string, productID uint) *SerialNumber {
	return &SerialNumber{
		Number:    number,
		ProductID: productID,
	}
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	category := NewCategory("Cozinha")
	db.Create(&category)

	product := NewProduct("Panela", 1500, category.ID)
	db.Create(&product)

	db.Create(NewSerialNumber("987654321", product.ID))

	// var products []Product
	// db.Preload("Category").Preload("SerialNumber").Find(&products)
	// for _, p := range products {
	// 	fmt.Println(p.Category.Name, "-", p.Name, "-", p.Price, "-", p.SerialNumber.Number)
	// }

	var categories []Category
	if err = db.Model(&Category{}).Preload("Products.SerialNumber").Find(&categories).Error; err != nil {
		panic(err)
	}

	for _, c := range categories {
		fmt.Println(c.Name, ":")
		for _, p := range c.Products {
			fmt.Println(" -", p.Name, "-", p.Price, "-", p.SerialNumber.Number)
		}
	}
}
