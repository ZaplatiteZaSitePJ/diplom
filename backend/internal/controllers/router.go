package controllers

import (
	"inno-accounting/internal/adapters/jwt"
	"inno-accounting/internal/controllers/handlers"
	"inno-accounting/internal/controllers/middleware"

	"github.com/gorilla/mux"
)

const urlPrefix = "/api/v1"

func InitRouter(h *handlers.Handlers, jwtManager jwt.TokenManager) *mux.Router {
	router := mux.NewRouter()

	// PUBLIC
	router.HandleFunc(urlPrefix+"/auth/register", h.CreateUser).Methods("POST")
	router.HandleFunc(urlPrefix+"/auth/login", h.Login).Methods("POST")

	// PROTECTED
	protected := router.PathPrefix(urlPrefix).Subrouter()
	protected.Use(middleware.JWTMiddleware(jwtManager))

	me := protected.PathPrefix("/me").Subrouter()
	me.HandleFunc("/profile", h.GetUserByMe).Methods("GET")

	// ADMIN
	admin := protected.PathPrefix("/admin").Subrouter()
	admin.Use(middleware.RoleMiddleware("admin"))

	admin.HandleFunc("/users", h.GetAllUsers).Methods("GET")       
	admin.HandleFunc("/users/{id}", h.GetUserByID).Methods("GET")

	admin.HandleFunc("/storages", h.GetAllStorages).Methods("GET")
	admin.HandleFunc("/storages/{id}", h.GetStorageByID).Methods("GET")
	admin.HandleFunc("/storages", h.CreateStorage).Methods("POST")
	admin.HandleFunc("/storages/{id}", h.UpdateStorageByID).Methods("PATCH")

	admin.HandleFunc("/items/tech", h.CreateTech).Methods("POST")
	admin.HandleFunc("/items/tech", h.GetAllTech).Methods("GET")
	
	return router
}