package user

import (
	"Bakers_backend/internal/entities"
	"Bakers_backend/internal/repository"
	"Bakers_backend/internal/service"
	"context"
	"golang.org/x/crypto/bcrypt"
)

type ServiceUser struct {
	UserRepo repository.UserRepo
}

func InitUserService(userRepo repository.UserRepo) service.UserServ {
	return &ServiceUser{UserRepo: userRepo}
}

func (usr ServiceUser) Create(ctx context.Context, user entities.UserCreate) (int, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	newUser := entities.UserCreate{
		UserBase: user.UserBase,
		Password: string(hashedPassword),
	}

	id, err := usr.Create(ctx, newUser)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (usr ServiceUser) GetUserByID(ctx context.Context, id int) (entities.User, error) {
	//TODO implement me
	panic("implement me")
}
