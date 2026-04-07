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


func (u *UserService) CreateUser(input *dto.CreateUser) (*domain.User, error){
	logger.Info("Creating new user: ", input)

	var hashed_password string

	// Validation password
	if err := crypt_password.ValidatePassword(input.Password); err != nil {
		logger.Info("service", err)
		return nil, app_errors.Unprocessable("password too weak", err)
	}

	// Hashing password
	hashed_password, err := crypt_password.EncryptPassword(input.Password) 
	if err != nil {
		logger.Info("service", err)
		return nil, app_errors.Internal("server unavailable now. Try again later", err)
	}
	
	newUser := domain.User{
		Email: input.Email,
		Username: input.Username,
		HashedPassword: hashed_password,
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Validation user
	if err := newUser.ValidateUser(); err != nil {
		logger.Info("service", err)
		return nil, app_errors.Unprocessable("Invalid username or email", err)
	}

	return u.repo.Save(&newUser)
}

func (u *UserService) FindUserByID(userID uuid.UUID) (*domain.User,  error){
	logger.Info("Trying to find user: ", userID)

	findedUser, err := u.repo.FindByID(userID)
	if err != nil {
	 	return nil,  err
	}

	return findedUser,  nil
}

func (u *UserService) FindAllUsers() ([]*domain.User, error) {
	logger.Info("Trying to find all users")
	findedUsers, err := u.repo.FindAll()

	if err != nil {
	 	wErr := custom_errors.New(err, 500)
	 	wErr.AddResponseData("Internal server error")
	 	wErr.AddLogData(err.Error())
		return nil, wErr
	}
	return findedUsers, nil
}

func (u *UserService) DeleteByID(userID int) error {
	return nil
}