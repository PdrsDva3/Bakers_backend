package service

import (
	"Bakers_backend/internal/entities"
	"context"
)

type UserServ interface {
	Create(ctx context.Context, user entities.UserCreate) (int, error)
	Get(ctx context.Context, id int) (*entities.User, error)
	Login(ctx context.Context, user entities.UserLogin) (*entities.User, error)
}

type AdminService interface {
	AdminCreate(ctx context.Context, adminCreate entities.AdminCreate) (int, error)
	//Login(ctx context.Context, adminLogin entities.AdminLogin) (int, error)
	//UpdatePassword(ctx context.Context, StudentID int, newPassword string) error
	//GetMe(ctx context.Context, studentID int) (entities.Admin, error)
	//Delete(ctx context.Context, studentID int) error
}
