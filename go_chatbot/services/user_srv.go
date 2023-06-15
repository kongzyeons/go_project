package services

import (
	"fmt"
	"go_chatbot/models"
	"go_chatbot/repository"
	"go_chatbot/utils"

	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(model models.UserCreate) error
	LoginUser(model models.UserLogin) (tokenString string, err error)
	GetUser(user_id string) (user repository.Users, err error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return userService{userRepo}
}

func (obj userService) CreateUser(model models.UserCreate) error {
	userCheck := repository.Users{
		Username: model.Username,
	}
	_, err := obj.userRepo.GetByID(userCheck)
	if err == nil {
		return fmt.Errorf("username used")
	}
	user := repository.Users{
		UserID:   "user_id-" + uuid.NewString(),
		Username: model.Username,
		Password: model.Password,
	}
	return obj.userRepo.Create(user)

}

func (obj userService) LoginUser(model models.UserLogin) (tokenString string, err error) {
	userCheck := repository.Users{
		Username: model.Username,
		Password: model.Password,
	}
	user, err := obj.userRepo.GetByID(userCheck)
	if err != nil {
		return tokenString, fmt.Errorf("invalid username or password")
	}
	token := utils.JWT_NewWithClaims(user)
	tokenString, err = token.SignedString([]byte(utils.SecretKey))
	if err != nil {
		return tokenString, fmt.Errorf("invalid token")
	}
	return tokenString, err
}

func (obj userService) GetUser(user_id string) (user repository.Users, err error) {
	userCheck := repository.Users{
		UserID: user_id,
	}
	user, err = obj.userRepo.GetByID(userCheck)
	if err != nil {
		return user, fmt.Errorf("invalid uesr_id")
	}
	return user, err
}
