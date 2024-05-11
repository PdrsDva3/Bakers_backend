package service

import (
	"Bakers_backend/internal/entities"
	"context"
)

type AdminService interface {
	AdminCreate(ctx context.Context, adminCreate entities.AdminCreate) (int, error)
	Login(ctx context.Context, adminLogin entities.AdminLogin) (int, error)
	ChangePassword(ctx context.Context, adminID int, newPWD string) error
	GetMe(ctx context.Context, studentID int) (*entities.Admin, error)
	Delete(ctx context.Context, adminID int) error
}

type BreadService interface {
	BreadCreate(ctx context.Context, breadCreate entities.BreadBase) (int, error)
	GetBread(ctx context.Context, breadID int) (*entities.Bread, error)
}
