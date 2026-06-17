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
	router.HandleFunc(urlPrefix+"/auth/refresh", h.Refresh).Methods("POST")
	router.HandleFunc("/auth/activate", h.Activate).Methods("GET")

	// PROTECTED
	protected := router.PathPrefix(urlPrefix).Subrouter()
	protected.Use(middleware.JWTMiddleware(jwtManager))

	me := protected.PathPrefix("/me").Subrouter()
	me.HandleFunc("/profile", h.GetUserByMe).Methods("GET")
	me.HandleFunc("/logout", h.Logout).Methods("POST")
	me.HandleFunc("/items/tech", h.GetMyTech).Methods("GET")
	me.HandleFunc("/items/docs", h.GetMyDocuments).Methods("GET")
	me.HandleFunc("/items/software", h.GetMySoftware).Methods("GET")
	
	// ADMIN
	admin := protected.PathPrefix("/admin").Subrouter()
	admin.Use(middleware.RoleMiddleware("admin"))

	admin.HandleFunc("/users", h.GetAllUsers).Methods("GET")       
	admin.HandleFunc("/users/{id}", h.GetUserByID).Methods("GET")
	admin.HandleFunc("/users/{id}", h.DeleteUserByID).Methods("DELETE")

	admin.HandleFunc("/storages", h.GetAllStorages).Methods("GET")
	admin.HandleFunc("/storages/{id}", h.GetStorageByID).Methods("GET")
	admin.HandleFunc("/storages", h.CreateStorage).Methods("POST")
	admin.HandleFunc("/storages/{id}", h.UpdateStorageByID).Methods("PATCH")
	admin.HandleFunc("/storages/{id}", h.DeleteStorageByID).Methods("DELETE")

	admin.HandleFunc("/categories/{type_index}", h.GetCategoriesByTypeID).Methods("GET")
	admin.HandleFunc("/items/{id}", h.DeleteTechByID).Methods("DELETE")

	admin.HandleFunc("/items/tech", h.CreateTech).Methods("POST")
	admin.HandleFunc("/items/tech", h.GetAllTech).Methods("GET")
	admin.HandleFunc("/items/tech/{id}", h.GetTechByID).Methods("GET")
	admin.HandleFunc("/items/tech/{id}", h.PatchTechByID).Methods("PATCH")

	admin.HandleFunc("/items/docs", h.CreateDocument).Methods("POST")
	admin.HandleFunc("/items/docs", h.GetAllDocuments).Methods("GET")
	admin.HandleFunc("/items/docs/{id}", h.GetDocumentByID).Methods("GET")
	admin.HandleFunc("/items/docs/{id}", h.PatchDocumentByID).Methods("PATCH")

	admin.HandleFunc("/items/software", h.CreateSoftware).Methods("POST")
	admin.HandleFunc("/items/software", h.GetAllSoftware).Methods("GET")
	admin.HandleFunc("/items/software/{id}", h.GetSoftwareByID).Methods("GET")
	admin.HandleFunc("/items/software/{id}", h.PatchSoftwareByID).Methods("PATCH")
	
	return router
}