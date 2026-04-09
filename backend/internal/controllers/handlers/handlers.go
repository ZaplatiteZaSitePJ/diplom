package handlers

import (
	"inno-accounting/internal/use-cases/auth"
	"inno-accounting/internal/use-cases/item/tech"
	storages "inno-accounting/internal/use-cases/storage"
	"inno-accounting/internal/use-cases/user"
)

type Handlers struct {
	User *user.UserService
	Auth *auth.AuthService
	Storage *storages.StorageService
	TechItems *tech.TechService
}

func New(user *user.UserService, auth *auth.AuthService, storage *storages.StorageService, tech *tech.TechService) *Handlers {
	return &Handlers{
		User: user,
		Auth: auth,
		Storage: storage,
		TechItems: tech,
	}
}