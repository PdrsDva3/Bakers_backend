package repository

import (
	"Bakers_backend/internal/entities"
	"context"
)

type UserRepo interface {
	CreateUser(ctx context.Context, create entities.UserCreate) (int, error)
	GetUser(ctx context.Context, id int) (entities.User, error)
	GetPasswordByPhone(ctx context.Context, phone string) (string, error)
	UpdatePasswordByID(ctx context.Context, id int, password string) error
	UpdateUserNameByID(ctx context.Context, id int, name string) error
	AddOrderByID(ctx context.Context, id int) error
}
