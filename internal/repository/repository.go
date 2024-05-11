package repository

import (
	"Bakers_backend/internal/entities"
	"context"
)

type AdminRepo interface {
	CreateAdmin(ctx context.Context, admin entities.AdminCreate) (int, error)
	GetAdminByID(ctx context.Context, adminID int) (*entities.Admin, error)
	GetPasswordByPhone(ctx context.Context, phone int64) (int, string, error)
	UpdatePasswordByID(ctx context.Context, adminID int, newPassword string) error
	DeleteAdmin(ctx context.Context, adminID int) error
}

type UserRepo interface {
	Create(ctx context.Context, create entities.UserCreate) (int, error)
	Get(ctx context.Context, id int) (*entities.User, error)
	GetHashedPasswordByPhone(ctx context.Context, phone int64) (int, string, error)
	UpdatePasswordByID(ctx context.Context, id int, password string) error
	UpdateNameByID(ctx context.Context, id int, name string) error
	AddOrderByID(ctx context.Context, id int) error
	DeleteByID(ctx context.Context, id int) error
	DeleteBreadByID(ctx context.Context, id int) error // хз что тут делать
}

type OrderRepo interface {
	Create(ctx context.Context, order entities.OrderBase) (int, error)
	GetByID(ctx context.Context, id int) (entities.Order, error)
	AddBreadByID(ctx context.Context, id int) error
	DeleteBreadByID(ctx context.Context, id int) error
	DeleteByID(ctx context.Context, id int) error
}

type BreadRepo interface {
	Create(ctx context.Context, bread entities.BreadBase) (int, error)
	GetBreadByID(ctx context.Context, breadID int) (*entities.Bread, error)
	ChangeCountBreadByID(ctx context.Context, breadID int, count int64) (int64, error)
	DeleteBread(ctx context.Context, breadID int) error
}
