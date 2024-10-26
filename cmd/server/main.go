package main

import (
	"net/http"

	"github.com/andreflor21/go-api/configs"
	"github.com/andreflor21/go-api/internal/entity"
	"github.com/andreflor21/go-api/internal/infra/database"
	"github.com/andreflor21/go-api/internal/infra/webserver/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main(){
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.User{}, &entity.Product{})
	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", configs.TokenAuth))
	r.Use(middleware.WithValue("jwtExpiresIn", configs.JwtExpiresIn))
	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(configs.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/", productHandler.GetProducts)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})
	r.Post("/users", userHandler.CreateUser)
	r.Post("/users/generate_jwt", userHandler.GetJWT)
	http.ListenAndServe(":8000", r)
}

// func LogRequest(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){
// 			log.Printf("Request: %s %s", r.Method, r.URL.Path)
// 			next.ServeHTTP(w, r)
// 		})
// }