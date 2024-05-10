package repository

import (
	"Bakers_backend/internal/entities"
	"context"
)

type UserRepo interface {
	Create(ctx context.Context, create entities.UserCreate) (int, error)
	Get(ctx context.Context, id int) (entities.User, error)
	GetPasswordByPhone(ctx context.Context, phone string) (string, error)
	UpdatePasswordByID(ctx context.Context, id int, password string) error
	UpdateNameByID(ctx context.Context, id int, name string) error
	AddOrderByID(ctx context.Context, id int) error
}

type OrderRepo interface {
	Create(ctx context.Context, order entities.OrderBase) (int, error)
	GetByID(ctx context.Context, id int) (entities.Order, error)
	AddBreadByID(ctx context.Context, id int) error
	DeleteBreadByID(ctx context.Context, id int) error
	DeleteByID(ctx context.Context, id int) error
}
