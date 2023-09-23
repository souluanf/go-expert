package main

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/souluanf/fullcycle-sqlc/internal/db"
)

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()
	queries := db.New(dbConn)
	// Create
	//err = queries.CreateCategory(ctx, db.CreateCategoryParams{
	//	ID:          uuid.New().String(),
	//	Name:        "Go",
	//	Description: sql.NullString{String: "Golang", Valid: true},
	//})
	//if err != nil {
	//	panic(err)
	//}

	// Get All
	//categories, err := queries.ListCategories(ctx)
	//if err != nil {
	//	panic(err)
	//}
	//
	//for _, category := range categories {
	//	println(category.ID, category.Name, category.Description.String)
	//}

	// Update
	err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
		ID:          "69c479a6-04ff-4f4c-8d4a-95a2463aba7a",
		Name:        "Go Expert",
		Description: sql.NullString{String: "Advanced Golang", Valid: true},
	})

	// Delete
	err = queries.DeleteCategory(ctx, "3d578fc7-ee6c-47ea-ba2d-104f752a186a")
	if err != nil {
		panic(err)
	}

	//Get All
	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		println(category.ID, category.Name, category.Description.String)
	}
}
