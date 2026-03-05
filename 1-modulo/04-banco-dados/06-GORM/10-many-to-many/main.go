package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID         uint       `gorm:"primarykey"`
	Name       string     `gorm:"size:255;not null"`
	Price      float64    `gorm:"not null"`
	Categories []Category `gorm:"many2many:products_categories;"`
	gorm.Model
}

type Category struct {
	ID       uint      `gorm:"primarykey"`
	Name     string    `gorm:"size:255;not null"`
	Products []Product `gorm:"many2many:products_categories;"`
	gorm.Model
}

func NewProduct(name string, price float64, categories []Category) *Product {
	return &Product{
		Name:       name,
		Price:      price,
		Categories: categories,
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

	// category := NewCategory("Cozinha")
	// db.Create(&category)

	// category2 := NewCategory("Eletrônico")
	// db.Create(&category2)

	// product := NewProduct("Panela", 1500, []Category{*category, *category2})
	// db.Create(&product)

	var categories []Category
	if err = db.Model(&Category{}).Preload("Products").Find(&categories).Error; err != nil {
		panic(err)
	}

	for _, c := range categories {
		fmt.Println(c.Name, ":")
		for _, p := range c.Products {
			fmt.Println(" -", p.Name, "-", p.Price)
		}
	}
}
