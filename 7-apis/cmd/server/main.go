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
	"log"
	"net/http"
)

// @title           Products API
// @version         1.0
// @description     Demo para aprendizado de APIs com Go
// @termsOfService  http://swagger.io/terms/

// @contact.name   Luan Fernandes
// @contact.url    https://luanfernandes.dev
// @contact.email  hello@luanfernandes.dev

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.apiKey  ApiKeyAuth
// @in  header
// @name  Authorization

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func main() {
	config, err := configs.LoadConfig(".")
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
	userHandler := handlers.NewUserHandler(userDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", config.TokenAuth))
	r.Use(middleware.WithValue("JwtExpiresIn", config.JWTExpiresIn))

	//r.Use(LogRequest)// Middleware criado por mim
	r.Route("/products", func(c chi.Router) {

		c.Use(jwtauth.Verifier(config.TokenAuth))
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

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})

}
