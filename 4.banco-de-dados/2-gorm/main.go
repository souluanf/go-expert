package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int     `gorm:"primary_key"`
	Name  string  `gorm:"type:varchar(100);not null"`
	Price float64 `gorm:"not null"`
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Product{})
	if err != nil {
		return
	}
	// create one
	//db.Create(&Product{Name: "Notebook", Price: 2000})

	// create many
	//db.Create(&[]Product{
	//	{Name: "Notebook", Price: 2000},
	//	{Name: "Mouse", Price: 20},
	//	{Name: "Keyboard", Price: 100},
	//})

	// select one
	//var product Product
	//db.First(&product, 1)
	//fmt.Println(product)
	//db.First(&product, "name = ?", "Mouse")
	//fmt.Println(product)

	// select all
	var products []Product
	db.Find(&products)
	for _, product := range products {
		fmt.Println(product)
	}
}
