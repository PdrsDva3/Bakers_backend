package service

import (
	"Bakers_beckend/internal/entities"
	"context"
)

type AdminService interface {
	AdminCreate(ctx context.Context, adminCreate entities.AdminCreate) (int, error)
	//Login(ctx context.Context, adminLogin entities.AdminLogin) (int, error)
	//UpdatePassword(ctx context.Context, StudentID int, newPassword string) error
	//GetMe(ctx context.Context, studentID int) (entities.Admin, error)
	//Delete(ctx context.Context, studentID int) error
}
