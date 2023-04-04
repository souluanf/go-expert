package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/souluanf/go-expert/7-apis/configs"
	"github.com/souluanf/go-expert/7-apis/internal/entity"
	"github.com/souluanf/go-expert/7-apis/internal/infra/database"
	"github.com/souluanf/go-expert/7-apis/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&entity.Product{}, &entity.User{})
	if err != nil {
		panic(err)
	}

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/products", productHandler.GetProducts)
	r.Post("/products", productHandler.CreateProduct)
	r.Get("/products/{id}", productHandler.GetProduct)
	r.Put("/products/{id}", productHandler.UpdateProduct)
	r.Delete("/products/{id}", productHandler.DeleteProduct)
	err = http.ListenAndServe("localhost:8080", r)
	if err != nil {
		panic(err)
	}
}
