package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Category has many Products
type Category struct {
	ID       int    `gorm:"primary_key"`
	Name     string `gorm:"type:varchar(100);not null"`
	Products []Product
}

// Product belongs to Category
// Product has one SerialNumber
type Product struct {
	ID           int     `gorm:"primary_key"`
	Name         string  `gorm:"type:varchar(100);not null"`
	Price        float64 `gorm:"not null"`
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int    `gorm:"primary_key"`
	Number    string `gorm:"type:varchar(100);not null"`
	ProductID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})
	if err != nil {
		return
	}

	// Insert category
	category := Category{Name: "Electronics"}
	db.Create(&category)

	// Insert product
	product := Product{Name: "Notebook", Price: 2000.00, CategoryID: category.ID}
	db.Create(&product)

	// Insert serial number
	serialNumber := SerialNumber{Number: "123456", ProductID: product.ID}
	db.Create(&serialNumber)

	// Select category and print products
	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		for _, product := range category.Products {
			fmt.Println(category.Name, product.Name, product.Price, product.SerialNumber.Number)
		}
	}

	// Select product
	var products []Product
	db.Preload("Category").Preload("SerialNumber").Find(&products)
	for _, product := range products {
		fmt.Println(product.Name, product.Price, product.Category.Name, product.SerialNumber.Number)
	}

}
