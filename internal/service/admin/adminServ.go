package adminserv

import (
	"Bakers_backend/internal/entities"
	"Bakers_backend/internal/repository"
	"Bakers_backend/internal/service"
	"context"
	"golang.org/x/crypto/bcrypt"
)

func InitAdminService(adminRepo repository.AdminRepo) service.AdminService {
	return AdminService{AdminRepo: adminRepo}
}

type AdminService struct {
	AdminRepo repository.AdminRepo
}

func (adm AdminService) AdminCreate(ctx context.Context, adminCreate entities.AdminCreate) (int, error) {
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(adminCreate.Password), 10)
	if err != nil {
		return 0, nil
	}

	newStudent := entities.AdminCreate{
		AdminBase: adminCreate.AdminBase,
		Password:  string(hashed_password),
	}

	id, err := adm.AdminRepo.CreateAdmin(ctx, newStudent)
	if err != nil {
		return 0, err
	}

	return id, nil
}
