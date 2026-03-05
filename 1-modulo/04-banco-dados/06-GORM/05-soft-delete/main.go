package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    uint    `gorm:"primarykey"`
	Name  string  `gorm:"size:255;not null"`
	Price float64 `gorm:"not null"`
	gorm.Model
}

func NewProduct(name string, price float64) *Product {
	return &Product{
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

	// db.Create(NewProduct("Notebook", 1500))

	// var p Product
	// db.First(&p, 1)
	// p.Name = "New Mouse"
	// db.Save(&p)

	// var p Product
	// db.First(&p, 1)
	// fmt.Println(p)
	// db.Delete(&p)

	var p Product
	db.First(&p, 1)
	fmt.Println(p)
}
