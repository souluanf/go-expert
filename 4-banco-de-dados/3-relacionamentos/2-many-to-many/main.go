package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Category has many Products
type Category struct {
	ID       int       `gorm:"primary_key"`
	Name     string    `gorm:"type:varchar(100);not null"`
	Products []Product `gorm:"many2many:products_categories;"`
}

// Product has many Category
type Product struct {
	ID         int        `gorm:"primary_key"`
	Name       string     `gorm:"type:varchar(100);not null"`
	Price      float64    `gorm:"not null"`
	Categories []Category `gorm:"many2many:products_categories;"`
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

	category1 := Category{Name: "Electronics"}
	category2 := Category{Name: "Books"}
	db.Create(&Product{Name: "Notebook", Price: 2000.00, Categories: []Category{category1, category2}})

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		for _, product := range category.Products {
			fmt.Printf("%s - %s\n", category.Name, product.Name)
		}
	}
}
