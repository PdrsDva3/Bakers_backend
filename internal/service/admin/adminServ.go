package adminserv

import (
	"Bakers_backend/internal/entities"
	"Bakers_backend/internal/repository"
	"Bakers_backend/internal/service"
	"Bakers_backend/pkg/customerr"
	"context"
	"golang.org/x/crypto/bcrypt"
)

func InitAdminService(adminRepo repository.AdminRepo) service.AdminService {
	return AdminService{AdminRepo: adminRepo}
}

type AdminService struct {
	AdminRepo repository.AdminRepo
}

func (adm AdminService) Login(ctx context.Context, adminLogin entities.AdminLogin) (int, error) {
	id, pwd, err := adm.AdminRepo.GetPasswordByPhone(ctx, adminLogin.Phone)
	if err != nil {
		return 0, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(pwd), []byte(adminLogin.Password))
	if err != nil {
		return 0, cerr.Err(cerr.Admin, cerr.Service, cerr.InvalidPWD, err).Error()
	}
	return id, nil
}

func (adm AdminService) GetMe(ctx context.Context, adminID int) (*entities.Admin, error) {
	admin, err := adm.AdminRepo.GetAdminByID(ctx, adminID)
	if err != nil {
		return nil, err
	}
	if admin == nil {
		return nil, cerr.Err(cerr.Admin, cerr.Service, cerr.NotFound, nil).Error()
	}
	return admin, nil
}

func (adm AdminService) AdminCreate(ctx context.Context, adminCreate entities.AdminCreate) (int, error) {
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(adminCreate.Password), 10)
	if err != nil {
		return 0, cerr.Err(cerr.Admin, cerr.Service, cerr.Hash, err).Error()
	}
	newAdmin := entities.AdminCreate{
		AdminBase: adminCreate.AdminBase,
		Password:  string(hashed_password),
	}

	id, err := adm.AdminRepo.CreateAdmin(ctx, newAdmin)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (adm AdminService) Delete(ctx context.Context, adminID int) error {
	err := adm.AdminRepo.DeleteAdmin(ctx, adminID)
	if err != nil {
		return err
	}
	return nil
}

func (adm AdminService) ChangePassword(ctx context.Context, adminID int, newPWD string) error {
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(newPWD), 10)
	if err != nil {
		return cerr.Err(cerr.Admin, cerr.Service, cerr.Hash, err).Error()
	}
	err = adm.AdminRepo.UpdatePasswordByID(ctx, adminID, string(hashed_password))
	if err != nil {
		return err
	}
	return nil
}
