package user

import (
	"Bakers_backend/internal/entities"
	"Bakers_backend/internal/repository"
	"Bakers_backend/pkg/customerr"
	"context"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func InitUserRepository(db *sqlx.DB) repository.UserRepo {
	return Repository{
		db}
}

func (user Repository) Create(ctx context.Context, create entities.UserCreate) (int, error) {
	var id int
	transaction, err := user.db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, customerr.ErrorMessage(0, err)
	}
	row := transaction.QueryRowContext(ctx, `INSERT INTO users (phone, name, hashed_password) VALUES ($1, $2, $3) returning id;`,
		create.Phone, create.Name, create.Password)

	err = row.Scan(&id)
	if err != nil {
		return 0, customerr.ErrorMessage(6, err)

	}

}

func (user Repository) Get(ctx context.Context, id int) (entities.User, error) {
	//TODO implement me
	panic("implement me")
}

func (user Repository) GetPasswordByPhone(ctx context.Context, phone string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (user Repository) UpdatePasswordByID(ctx context.Context, id int, password string) error {
	//TODO implement me
	panic("implement me")
}

func (user Repository) UpdateNameByID(ctx context.Context, id int, name string) error {
	//TODO implement me
	panic("implement me")
}

func (user Repository) AddOrderByID(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}
