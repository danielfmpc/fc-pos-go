package main

import (
	"log"
	"net/http"
	"pos-go-api/configs"
	"pos-go-api/internal/entity"
	"pos-go-api/internal/infra/database"
	"pos-go-api/internal/infra/webserver/handlers"

	_ "pos-go-api/docs"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title Go API Product
// @version 1.0
// @description Product API with JWT Authentication
// @termsOfService http://swagger.io/terms/

// @contact.name Daniel
// @contact.email daniel@daniel.com
// @contact.url https://daniel.com

// @license.name Daniel
// @license.url https://daniel.com

// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("teste.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDb := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDb)

	userDb := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDb)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", config.TokenAuth))
	r.Use(middleware.WithValue("jwtExpiration", config.JWTExpiration))
	// r.Use(LogRequest)
	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(config.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.Create)
		r.Get("/{id}", productHandler.GetProductByID)
		r.Get("/", productHandler.GetProducts)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Post("/users", userHandler.Create)
	r.Post("/users/generate_token", userHandler.GetJWT)

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/docs/doc.json")))

	http.ListenAndServe(":8080", r)
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request: ", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
