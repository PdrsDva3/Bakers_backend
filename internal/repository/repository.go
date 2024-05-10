package repository

import (
	"Bakers_beckend/internal/entities"
	"context"
)

type AdminRepo interface {
	CreateAdmin(ctx context.Context, admin entities.AdminCreate) (int, error)
	GetAdminByID(ctx context.Context, adminID int) (*entities.Admin, error)
	GetPasswordByPhone(ctx context.Context, phone int) (string, error)
	UpdatePasswordByID(ctx context.Context, adminID int, newPassword string) error
	DeleteAdmin(ctx context.Context, adminID int) error
	CreateBread(ctx context.Context, bread entities.BreadBase) (int, error)
}
