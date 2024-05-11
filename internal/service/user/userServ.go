package user

import (
	"Bakers_backend/internal/entities"
	"Bakers_backend/internal/repository"
	"Bakers_backend/internal/service"
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"strconv"
	//"regexp"
)

type ServiceUser struct {
	UserRepo repository.UserRepo
}

func InitUserService(userRepo repository.UserRepo) service.UserServ {
	return &ServiceUser{UserRepo: userRepo}
}

//todo валидация почты))
//func validateEmail(email string) bool {
//	// Регулярное выражение для проверки email
//	regex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
//	return regex.MatchString(email)
//}

func validatePhone(phone string) bool {
	// Регулярное выражение для проверки 10 цифр российского номера телефона
	regex := regexp.MustCompile(`^\d{10}$`)
	return regex.MatchString(phone)

}

func (usr ServiceUser) Create(ctx context.Context, user entities.UserCreate) (int, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	newUser := entities.UserCreate{
		UserBase: user.UserBase,
		Password: string(hashedPassword),
	}

	flag := validatePhone(strconv.FormatInt(user.Phone, 10))
	if !flag {
		return 0, errors.New("invalid phone")
	}

	id, err := usr.UserRepo.Create(ctx, newUser)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (usr ServiceUser) Get(ctx context.Context, id int) (*entities.User, error) {
	user, err := usr.UserRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("User not found")
	}

	return user, nil
}

func (usr ServiceUser) Login(ctx context.Context, user entities.UserLogin) (*entities.User, error) {
	id, hashed_password, err := usr.UserRepo.GetHashedPasswordByPhone(ctx, user.Phone)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashed_password), []byte(user.Password))
	if err != nil {
		return nil, errors.New("Invalid password")
	}
	userr, err := usr.UserRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return userr, nil
}
