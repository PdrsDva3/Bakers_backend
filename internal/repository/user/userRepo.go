package user

import (
	"Bakers_backend/internal/entities"
	"Bakers_backend/internal/repository"
	"Bakers_backend/pkg/customerr"
	"context"
	"github.com/jmoiron/sqlx"
)

type RepositoryUser struct {
	db *sqlx.DB
}

func InitUserRepository(db *sqlx.DB) repository.UserRepo {
	return RepositoryUser{
		db}
}

func (user RepositoryUser) Create(ctx context.Context, create entities.UserCreate) (int, error) {
	var id int
	transaction, err := user.db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, customerr.ErrorMessage(0, err)
	}
	row := transaction.QueryRowContext(ctx, `INSERT INTO users (phone, name, hashed_password) VALUES ($1, $2, $3) returning id;`,
		create.Phone, create.Name, create.Password)

	err = row.Scan(&id)
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return 0, customerr.ErrorMessage(3, err)
		}
		return 0, customerr.ErrorMessage(6, err)
	}
	if err := transaction.Commit(); err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return 0, customerr.ErrorMessage(3, err)
		}
		return 0, customerr.ErrorMessage(5, err)
	}
	return id, nil
}

func (user RepositoryUser) Get(ctx context.Context, id int) (entities.User, error) {
	var OldUser entities.User
	rows := user.db.QueryRowContext(ctx, `SELECT phone, name FROM users WHERE id = $1;`, id)

	err := rows.Scan(&OldUser.Phone, &OldUser.Name)
	if err != nil {
		return entities.User{}, customerr.ErrorMessage(6, err)
	}
	OldUser.ID = id
	return OldUser, nil
}

func (user RepositoryUser) GetHashedPasswordByPhone(ctx context.Context, phone string) (string, error) {
	var hsh_password string
	row := user.db.QueryRowContext(ctx, `SELECT hashed_password FROM users WHERE phone = $1;`, phone)
	err := row.Scan(&hsh_password)
	if err != nil {
		return "", customerr.ErrorMessage(6, err)
	}
	return hsh_password, nil
}

func (user RepositoryUser) UpdatePasswordByID(ctx context.Context, id int, password string) error {
	transaction, err := user.db.BeginTxx(ctx, nil)
	if err != nil {
		return customerr.ErrorMessage(0, err)
	}
	result, err := transaction.ExecContext(ctx, `UPDATE users SET hached_password=$2 WHERE id=$1;`, id, password)
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return customerr.ErrorMessage(3, err)
		}
		return customerr.ErrorMessage(7, err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return customerr.ErrorMessage(3, err)
		}
		return customerr.ErrorMessage(7, err)
	}
	if count != 1 {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return customerr.ErrorMessage(3, err)
		}
		return customerr.ErrorMessage(8, err)
	}

	if err := transaction.Commit(); err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return customerr.ErrorMessage(3, err)
		}
		return customerr.ErrorMessage(5, err)
	}
	return nil
}

func (user RepositoryUser) UpdateNameByID(ctx context.Context, id int, name string) error {
	transaction, err := user.db.BeginTxx(ctx, nil)
	if err != nil {
		return customerr.ErrorMessage(0, err)
	}
	result, err := transaction.ExecContext(ctx, `UPDATE users SET name=$2 WHERE id=$1;`, id, name)
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return customerr.ErrorMessage(3, err)
		}
		return customerr.ErrorMessage(7, err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return customerr.ErrorMessage(3, err)
		}
		return customerr.ErrorMessage(7, err)
	}
	if count != 1 {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return customerr.ErrorMessage(3, err)
		}
		return customerr.ErrorMessage(8, err)
	}

	if err := transaction.Commit(); err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return customerr.ErrorMessage(3, err)
		}
		return customerr.ErrorMessage(5, err)
	}
	return nil
}

func (user RepositoryUser) DeleteByID(ctx context.Context, id int) error {
	transaction, err := user.db.BeginTxx(ctx, nil)
	if err != nil {
		return customerr.ErrorMessage(0, err)
	}
	result, err := transaction.ExecContext(ctx, `DELETE FROM users WHERE id=$1;`, id)
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return customerr.ErrorMessage(3, err)
		}
		return customerr.ErrorMessage(7, err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return customerr.ErrorMessage(3, err)
		}
		return customerr.ErrorMessage(7, err)
	}
	if count != 1 {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return customerr.ErrorMessage(3, err)
		}
		return customerr.ErrorMessage(8, err)
	}

	if err := transaction.Commit(); err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return customerr.ErrorMessage(3, err)
		}
		return customerr.ErrorMessage(5, err)
	}
	return nil

}

func (user RepositoryUser) DeleteBreadByID(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func (user RepositoryUser) AddOrderByID(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}
