package adminserv

import (
	"Bakers_beckend/internal/entities"
	"Bakers_beckend/internal/repository"
	"Bakers_beckend/internal/service"
	"context"
)

type AdminService struct {
	AdminRepo repository.AdminRepo
}

func (adm AdminService) AdminCreate(ctx context.Context, adminCreate entities.AdminCreate) (int, error) {
	//hashed_password, err := bcrypt.GenerateFromPassword([]byte(student.Password), 10)
	//if err != nil {
	//	return 0, nil
	//}
	//
	//newStudent := entities.CreateStudent{
	//	StudentBase: student.StudentBase,
	//	Password:    string(hashed_password),
	//}
	//
	//id, err := usrs.StudentRepository.Create(ctx, newStudent)
	//if err != nil {
	//	return 0, err
	//}
	//
	//return id, nil
	panic("aaaaa")
}

func InitAdminService(adminRepo repository.AdminRepo) service.AdminService {
	return AdminService{AdminRepo: adminRepo}
}
