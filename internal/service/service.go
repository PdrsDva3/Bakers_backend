package service

import (
	"Bakers_backend/internal/entities"
	"context"
)

type UserServ interface {
	Create(ctx context.Context, user entities.UserCreate) (int, error)
	Get(ctx context.Context, id int) (*entities.User, error)
	Login(ctx context.Context, user entities.UserLogin) (*entities.User, error)
	ChangePassword(ctx context.Context, id int, newPassword string) error
	ChangeName(ctx context.Context, id int, name string) error
	Delete(ctx context.Context, id int) error
}

type AdminService interface {
	AdminCreate(ctx context.Context, adminCreate entities.AdminCreate) (int, error)
	Login(ctx context.Context, adminLogin entities.AdminLogin) (int, error)
	ChangePassword(ctx context.Context, adminID int, newPWD string) error
	GetMe(ctx context.Context, adminID int) (*entities.Admin, error)
	Delete(ctx context.Context, adminID int) error
}

type BreadService interface {
	BreadCreate(ctx context.Context, breadCreate entities.BreadBase) (int, error)
	GetBread(ctx context.Context, breadID int) (*entities.Bread, error)
	ChangeBread(ctx context.Context, breadID int, count int) (int, error)
	DeleteBread(ctx context.Context, breadID int) error
}
