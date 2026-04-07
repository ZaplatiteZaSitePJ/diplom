package dto

import (
	"inno-accounting/internal/domain"
)

type CreateUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type PublicUser struct {
	Username         string   `json:"username"`
	Email            string   `json:"email"`
	Role             string   `json:"role"`
}

func PublicUserFromModel(user *domain.User) *PublicUser {
	return &PublicUser{
		Username: user.Username, 
		Email: user.Email,
		Role: user.Role,
	}
}

func SeveralUsersToPublic(users []*domain.User) []*PublicUser {
    publicUsers := make([]*PublicUser, 0, len(users))
    for _, u := range users {
        publicUsers = append(publicUsers, PublicUserFromModel(u))
    }
    return publicUsers
}