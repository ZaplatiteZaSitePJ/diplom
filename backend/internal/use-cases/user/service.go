package user

import (
	"inno-accounting/internal/domain"
	"inno-accounting/internal/dto"
	"inno-accounting/pkg/crypt_password"
	"inno-accounting/pkg/logger"
	"inno-accounting/pkg/server_utils/app_errors"
	custom_errors "inno-accounting/pkg/server_utils/errors"
	"time"

	"github.com/google/uuid"
)

type UserService struct {
	repo UserRepository
}

func New(repo UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

//
// 🔹 CREATE USER
//
func (s *UserService) CreateUser(input *dto.CreateUser) (*domain.User, error) {
	logger.Info("Creating new user: ", input)

	if err := crypt_password.ValidatePassword(input.Password); err != nil {
		return nil, app_errors.Unprocessable("password too weak", err)
	}

	hashedPassword, err := crypt_password.EncryptPassword(input.Password)
	if err != nil {
		return nil, app_errors.Internal("server unavailable now. Try again later", err)
	}

	user := &domain.User{
		ID:             uuid.New(),
		Email:          input.Email,
		Name:           input.Name,
		LastName:       input.LastName,
		Post:           input.Post,
		Grade:          input.Grade,
		City:           input.City,
		Role:           input.Role,
		HashedPassword: hashedPassword,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	if err := user.ValidateUser(); err != nil {
		return nil, app_errors.Unprocessable("invalid user data", err)
	}

	return s.repo.Save(user)
}

//
// 🔹 GET ALL USERS
//
func (s *UserService) FindAllUsers(filter *dto.UserFilter) ([]*domain.User, error) {
	logger.Info("Trying to find users with filter")

	users, err := s.repo.FindAll(filter)
	if err != nil {
		wErr := custom_errors.New(err, 500)
		wErr.AddResponseData("Internal server error")
		wErr.AddLogData(err.Error())
		return nil, wErr
	}

	return users, nil
}

//
// 🔹 GET USER BY ID
//
func (s *UserService) FindUserByID(userID uuid.UUID) (*domain.User, error) {
	logger.Info("Trying to find user: ", userID)

	user, err := s.repo.FindByID(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

//
// 🔹 UPDATE USER (PATCH)
//
func (s *UserService) UpdateUser(userID uuid.UUID, input *dto.UpdateUser) (*domain.User, error) {
	logger.Info("Updating user: ", userID)

	user, err := s.repo.FindByID(userID)
	if err != nil {
		return nil, err
	}

	if input.Email != nil {
		user.Email = *input.Email
	}
	if input.Name != nil {
		user.Name = *input.Name
	}
	if input.LastName != nil {
		user.LastName = *input.LastName
	}
	if input.Post != nil {
		user.Post = *input.Post
	}
	if input.Grade != nil {
		user.Grade = *input.Grade
	}
	if input.City != nil {
		user.City = *input.City
	}

	user.UpdatedAt = time.Now()

	if err := user.ValidateUser(); err != nil {
		return nil, app_errors.Unprocessable("invalid user data", err)
	}

	return s.repo.Update(user)
}

//
// 🔹 DELETE USER (SOFT DELETE)
//
func (s *UserService) DeleteByID(userID uuid.UUID) error {
	logger.Info("Deleting user: ", userID)

	return s.repo.Delete(userID)
}

//
// 🔹 FIND USER BY EMAIL
//
func (s *UserService) FindUserByEmail(email string) (*domain.User, error) {
	logger.Info("Trying to find user by email: ", email)

	if email == "" {
		return nil, app_errors.InvalidInput("email parameter is required", nil)
	}

	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}