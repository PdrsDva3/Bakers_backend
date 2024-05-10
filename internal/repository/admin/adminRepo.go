package admin

import (
	"Bakers_beckend/internal/entities"
	"Bakers_beckend/internal/repository"
	"context"
	"errors"
	"github.com/jmoiron/sqlx"
)

type AdminRepo struct {
	db *sqlx.DB
}

func (adm AdminRepo) CreateAdmin(ctx context.Context, admin entities.AdminCreate) (int, error) {
	var id int
	transaction, err := adm.db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, err
	}
	row := transaction.QueryRowContext(ctx, `insert into admin (phone, hashed_password) values ($1, $2) returning id;`,
		admin.Phone, admin.Password)
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}
	if err := transaction.Commit(); err != nil {
		return 0, err
	}
	return id, nil
}

func (adm AdminRepo) GetAdminByID(ctx context.Context, adminID int) (*entities.Admin, error) {
	var user entities.Admin
	row := adm.db.QueryRowContext(ctx, `select phone from admin where id = $1`, adminID)
	err := row.Scan(&user.Phone)
	user.AdminID = adminID
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (adm AdminRepo) GetPasswordByPhone(ctx context.Context, phone int) (string, error) {
	var password string
	row := adm.db.QueryRowContext(ctx, `select hashed_password from admin where phone = $1`, phone)
	err := row.Scan(&password)
	if err != nil {
		return "", err
	}
	return password, nil
}

func (adm AdminRepo) UpdatePasswordByID(ctx context.Context, adminID int, newPassword string) error {
	transaction, err := adm.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	result, err := transaction.ExecContext(ctx, `UPDATE admin SET hashed_password = $2 WHERE id = $1;`, adminID, newPassword)
	if err != nil {
		return err
	}

	count, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if count != 1 {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return errors.New("failed rollback")
		}
		return errors.New("failed to update password")
	}

	if err = transaction.Commit(); err != nil {
		return err
	}
	return nil
}

func (adm AdminRepo) DeleteAdmin(ctx context.Context, adminID int) error {
	transaction, err := adm.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	result, err := transaction.ExecContext(ctx, `DELETE FROM admin WHERE id=$1;`, adminID)
	if err != nil {
		return err
	}
	count, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if count != 1 {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return errors.New("failed rollback")
		}
		return errors.New("failed to update password")
	}
	if err = transaction.Commit(); err != nil {
		return err
	}

	return nil
}

func (adm AdminRepo) CreateBread(ctx context.Context, bread entities.BreadBase) (int, error) {
	var id int
	transaction, err := adm.db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, err
	}
	row := transaction.QueryRowContext(ctx, `insert into bread (name, price, description, count, photo) values ($1, $2, $3, $4, $5) returning id;`,
		bread.Name, bread.Price, bread.Description, bread.Count, bread.Photo)
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}
	if err := transaction.Commit(); err != nil {
		return 0, err
	}
	return id, nil
}

func NewAdminRepo(db *sqlx.DB) repository.AdminRepo {
	return AdminRepo{
		db: db,
	}
}
