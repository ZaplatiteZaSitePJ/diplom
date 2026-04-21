package dto

import (
	"inno-accounting/internal/domain"

	"github.com/google/uuid"
)

type CreateUser struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	Post     string `json:"post"`
	Grade    string `json:"grade"`
	City     string `json:"city"`
	Role     string `json:"role"`
	Password string `json:"password"` 
}

type UpdateUser struct {
	Email    *string `json:"email"`
	Name     *string `json:"name"`
	LastName *string `json:"lastname"`
	Post     *string `json:"post"`
	Grade    *string `json:"grade"`
	City     *string `json:"city"`
}

type PublicUser struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	LastName string    `json:"lastname"`
	Email    string    `json:"email"`
	Post     string    `json:"post"`
	Grade    string    `json:"grade"`
	City     string    `json:"city"`
}

type UserWithItems struct {
	PublicUser
	Items []domain.Item `json:"items"`
}

type UserFilter struct {
	ID       *string
	Email    *string
	Name     *string
	LastName *string
	Post     *string
	Grade    *string
	City     *string
}

func PublicUserFromModel(user *domain.User) *PublicUser {
	return &PublicUser{
		ID:       user.ID,
		Name:     user.Name,
		LastName: user.LastName,
		Email:    user.Email,
		Post:     user.Post,
		Grade:    user.Grade,
		City:     user.City,
	}
}

func SeveralUsersToPublic(users []*domain.User) []*PublicUser {
    publicUsers := make([]*PublicUser, 0, len(users))
    for _, u := range users {
        publicUsers = append(publicUsers, PublicUserFromModel(u))
    }
    return publicUsers
}