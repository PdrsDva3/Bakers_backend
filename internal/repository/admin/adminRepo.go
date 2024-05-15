package admin

import (
	"Bakers_backend/internal/entities"
	"Bakers_backend/internal/repository"
	"Bakers_backend/pkg/customerr"
	"context"
	"github.com/jmoiron/sqlx"
)

type AdminRepo struct {
	db *sqlx.DB
}

func InitAdminRepo(db *sqlx.DB) repository.AdminRepo {
	return AdminRepo{
		db: db,
	}
}

func (adm AdminRepo) CreateAdmin(ctx context.Context, admin entities.AdminCreate) (int, error) {
	var id int
	transaction, err := adm.db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, cerr.Err(cerr.Admin, cerr.Repository, cerr.Transaction, err).Error()
	}
	row := transaction.QueryRowContext(ctx, `insert into admin (phone, hashed_password) values ($1, $2) returning id;`,
		admin.Phone, admin.Password)
	err = row.Scan(&id)
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return 0, cerr.Err(cerr.Admin, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return 0, cerr.Err(cerr.Admin, cerr.Repository, cerr.Scan, err).Error()
	}
	if err := transaction.Commit(); err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return 0, cerr.Err(cerr.Admin, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return 0, cerr.Err(cerr.Admin, cerr.Repository, cerr.Commit, err).Error()
	}
	return id, nil
}

func (adm AdminRepo) GetAdminByID(ctx context.Context, adminID int) (*entities.Admin, error) {
	var user entities.Admin
	row := adm.db.QueryRowContext(ctx, `select phone from admin where id = $1`, adminID)
	err := row.Scan(&user.Phone)
	user.AdminID = adminID
	if err != nil {
		return nil, cerr.Err(cerr.Admin, cerr.Repository, cerr.Scan, err).Error()
	}
	return &user, nil
}

func (adm AdminRepo) GetPasswordByPhone(ctx context.Context, phone int64) (int, string, error) {
	var password string
	var id int
	row := adm.db.QueryRowContext(ctx, `select hashed_password, id from admin where phone = $1`, phone)
	err := row.Scan(&password, &id)
	if err != nil {
		return 0, "", cerr.Err(cerr.Admin, cerr.Repository, cerr.Scan, err).Error()
	}
	return id, password, nil
}

func (adm AdminRepo) UpdatePasswordByID(ctx context.Context, adminID int, newPassword string) error {
	transaction, err := adm.db.BeginTx(ctx, nil)
	if err != nil {
		return cerr.Err(cerr.Admin, cerr.Repository, cerr.Transaction, err).Error()
	}
	result, err := transaction.ExecContext(ctx, `UPDATE admin SET hashed_password = $2 WHERE id = $1;`, adminID, newPassword)
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Admin, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Admin, cerr.Repository, cerr.ExecCon, err).Error()
	}

	count, err := result.RowsAffected()
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Admin, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Admin, cerr.Repository, cerr.Rows, err).Error()
	}

	if count != 1 {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Admin, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Admin, cerr.Repository, cerr.NoOneRow, err).Error()
	}

	if err = transaction.Commit(); err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Admin, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Admin, cerr.Repository, cerr.Commit, err).Error()
	}
	return nil
}

func (adm AdminRepo) DeleteAdmin(ctx context.Context, adminID int) error {
	transaction, err := adm.db.BeginTx(ctx, nil)
	if err != nil {
		return cerr.Err(cerr.Admin, cerr.Repository, cerr.Transaction, err).Error()
	}

	result, err := transaction.ExecContext(ctx, `DELETE FROM admin WHERE id=$1;`, adminID)
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Admin, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Admin, cerr.Repository, cerr.ExecCon, err).Error()
	}
	count, err := result.RowsAffected()
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Admin, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Admin, cerr.Repository, cerr.Rows, err).Error()
	}
	if count != 1 {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Admin, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Admin, cerr.Repository, cerr.NoOneRow, err).Error()
	}
	if err = transaction.Commit(); err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Admin, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Admin, cerr.Repository, cerr.Commit, err).Error()
	}

	return nil
}
