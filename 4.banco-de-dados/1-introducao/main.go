package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Produto struct {
	ID    string
	Name  string
	Price float64
}

func NewProduto(name string, price float64) *Produto {
	return &Produto{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	product := NewProduto("Notebook", 1500.00)
	err = insertProduto(db, product)
	if err != nil {
		panic(err)
	}
}

func insertProduto(db *sql.DB, produto *Produto) error {
	stmt, err := db.Prepare("insert into produtos(id, name, price) values(?,?,?)")
	if err != nil {
		return nil
	}
	defer stmt.Close()
	_, err = stmt.Exec(produto.ID, produto.Name, produto.Price)
	if err != nil {
		return err
	}
	return nil
}
