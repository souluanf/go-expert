package main

import (
	"database/sql"
	"fmt"
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
	product.Price = 2000.00
	err = updateProduto(db, product)
	if err != nil {
		panic(err)
	}
	produto, err := selectProdut(db, product.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println(produto.Name, produto.Price)

	products, err := selectAllProdutos(db)
	if err != nil {
		panic(err)
	}
	for _, p := range products {
		fmt.Println(p.Name, p.Price)
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

func updateProduto(db *sql.DB, produto *Produto) error {
	stmt, err := db.Prepare("update produtos set name = ?, price = ? where id = ?")
	if err != nil {
		return nil
	}
	defer stmt.Close()
	_, err = stmt.Exec(produto.Name, produto.Price, produto.ID)
	if err != nil {
		return err
	}
	return nil
}

func selectProdut(db *sql.DB, id string) (*Produto, error) {
	stmt, err := db.Prepare("select * from produtos where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var produto Produto
	err = stmt.QueryRow(id).Scan(&produto.ID, &produto.Name, &produto.Price)
	if err != nil {
		return nil, err
	}
	return &produto, nil
}

func selectAllProdutos(db *sql.DB) ([]Produto, error) {
	stmt, err := db.Query("select id, name, price from produtos")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var products []Produto
	for stmt.Next() {
		var produto Produto
		err = stmt.Scan(&produto.ID, &produto.Name, &produto.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, produto)
	}
	return products, nil
}
