package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int    `gorm:"primary_key"`
	Name string `gorm:"type:varchar(100);not null"`
}

type Product struct {
	ID         int     `gorm:"primary_key"`
	Name       string  `gorm:"type:varchar(100);not null"`
	Price      float64 `gorm:"not null"`
	CategoryID int
	Category   Category
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Product{}, &Category{})
	if err != nil {
		return
	}

	// Insert category
	//category := Category{Name: "Electronics"}
	//db.Create(&category)

	// Insert product
	//product := Product{Name: "Notebook", Price: 2000.00, CategoryID: category.ID}
	//db.Create(&product)

	// Select product
	var products []Product
	db.Preload("Category").Find(&products)
	for _, product := range products {
		fmt.Println(product.Name, product.Price, product.Category.Name)
	}

}
