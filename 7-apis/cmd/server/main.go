package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/souluanf/go-expert/7-apis/configs"
	"github.com/souluanf/go-expert/7-apis/internal/entity"
	"github.com/souluanf/go-expert/7-apis/internal/infra/database"
	"github.com/souluanf/go-expert/7-apis/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

func main() {
	configs, err := configs.LoadConfig(".")
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

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB, configs.TokenAuth, configs.JWTExpiresIn)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/products", func(c chi.Router) {
		c.Use(jwtauth.Verifier(configs.TokenAuth))
		c.Use(jwtauth.Authenticator)
		c.Get("/", productHandler.GetProducts)
		c.Post("/", productHandler.CreateProduct)
		c.Get("/{id}", productHandler.GetProduct)
		c.Put("/{id}", productHandler.UpdateProduct)
		c.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Route("/users", func(c chi.Router) {
		c.Post("/", userHandler.Create)
		c.Post("/generate_token", userHandler.GetJWT)
	})

	err = http.ListenAndServe("localhost:8080", r)
	if err != nil {
		panic(err)
	}
}
