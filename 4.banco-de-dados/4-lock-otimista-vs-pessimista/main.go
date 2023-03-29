package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	db.AutoMigrate(&Product{}, &Category{})

	// create many
	category1 := Category{Name: "Electronics"}
	category2 := Category{Name: "Books"}
	db.Create(&Product{Name: "Notebook", Price: 2000.00, Categories: []Category{category1, category2}})

	tx := db.Begin()
	var c Category
	err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&c, 1).Error
	if err != nil {
		panic(err)
	}
	c.Name = "FOR UPDATE"
	err = tx.Debug().Save(&c).Error
	if err != nil {
		panic(err)
	}
	tx.Commit()

}
