package handlers

import (
	"inno-accounting/internal/use-cases/auth"
	storages "inno-accounting/internal/use-cases/storage"
	"inno-accounting/internal/use-cases/user"
)

type Handlers struct {
	User *user.UserService
	Auth *auth.AuthService
	Storage *storages.StorageService
}

func New(user *user.UserService, auth *auth.AuthService, storage *storages.StorageService) *Handlers {
	return &Handlers{
		User: user,
		Auth: auth,
		Storage: storage,
	}
}