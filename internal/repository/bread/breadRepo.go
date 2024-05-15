package bread

import (
	"Bakers_backend/internal/entities"
	"Bakers_backend/internal/repository"
	"Bakers_backend/pkg/customerr"
	"context"
	"github.com/jmoiron/sqlx"
)

type BreadRepo struct {
	db *sqlx.DB
}

func InitBreadRepo(db *sqlx.DB) repository.BreadRepo {
	return BreadRepo{
		db: db,
	}
}

func (brd BreadRepo) Create(ctx context.Context, bread entities.BreadBase) (int, error) {
	var id int
	transaction, err := brd.db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, cerr.Err(cerr.Bread, cerr.Repository, cerr.Transaction, err).Error()
	}
	row := transaction.QueryRowContext(ctx, `insert into bread (name, price, description, count, photo) values ($1, $2, $3, $4, $5) returning id;`,
		bread.Name, bread.Price, bread.Description, bread.Count, bread.Photo)
	err = row.Scan(&id)
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return 0, cerr.Err(cerr.Bread, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return 0, cerr.Err(cerr.Bread, cerr.Repository, cerr.Scan, err).Error()
	}
	if err := transaction.Commit(); err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return 0, cerr.Err(cerr.Bread, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return 0, cerr.Err(cerr.Bread, cerr.Repository, cerr.Commit, err).Error()
	}
	return id, nil
}

func (brd BreadRepo) GetBreadByID(ctx context.Context, breadID int) (*entities.Bread, error) {
	var bread entities.Bread
	row := brd.db.QueryRowContext(ctx, `select name, price, description, count, photo from bread where id = $1`, breadID)
	err := row.Scan(&bread.Name, &bread.Price, &bread.Description, &bread.Count, &bread.Photo)
	bread.ID = breadID
	if err != nil {
		return nil, cerr.Err(cerr.Bread, cerr.Repository, cerr.Scan, err).Error()
	}
	return &bread, nil
}

func (brd BreadRepo) DeleteBread(ctx context.Context, breadID int) error {
	transaction, err := brd.db.BeginTx(ctx, nil)
	if err != nil {
		return cerr.Err(cerr.Bread, cerr.Repository, cerr.Transaction, err).Error()
	}

	result, err := transaction.ExecContext(ctx, `DELETE FROM bread WHERE id=$1;`, breadID)
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Bread, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Bread, cerr.Repository, cerr.ExecCon, err).Error()
	}
	count, err := result.RowsAffected()
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Bread, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Bread, cerr.Repository, cerr.Rows, err).Error()
	}
	if count != 1 {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Bread, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Bread, cerr.Repository, cerr.NoOneRow, err).Error()
	}
	if err = transaction.Commit(); err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Bread, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Bread, cerr.Repository, cerr.Commit, err).Error()
	}

	return nil
}

func (brd BreadRepo) ChangeCountBreadByID(ctx context.Context, breadID int, cnt int) (int, error) {
	var oldCount int
	row := brd.db.QueryRowContext(ctx, `select (count) from bread where id = $1`, breadID)
	err := row.Scan(&oldCount)
	if err != nil {
		return 0, cerr.Err(cerr.Bread, cerr.Repository, cerr.Scan, err).Error()
	}
	transaction, err := brd.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, cerr.Err(cerr.Bread, cerr.Repository, cerr.Transaction, err).Error()
	}
	if oldCount+cnt < 0 {
		return 0, cerr.Err(cerr.Bread, cerr.Repository, cerr.InvalidCount, nil).Error()
	}
	result, err := transaction.ExecContext(ctx, `UPDATE bread SET count = $2 WHERE id = $1;`, breadID, cnt+oldCount)
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return 0, cerr.Err(cerr.Bread, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return 0, cerr.Err(cerr.Bread, cerr.Repository, cerr.ExecCon, err).Error()
	}

	count, err := result.RowsAffected()
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return 0, cerr.Err(cerr.Bread, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return 0, cerr.Err(cerr.Bread, cerr.Repository, cerr.Rows, err).Error()
	}

	if count != 1 {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return 0, cerr.Err(cerr.Bread, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return 0, cerr.Err(cerr.Bread, cerr.Repository, cerr.NoOneRow, err).Error()
	}

	if err = transaction.Commit(); err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return 0, cerr.Err(cerr.Bread, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return 0, cerr.Err(cerr.Bread, cerr.Repository, cerr.Commit, err).Error()
	}
	return cnt + oldCount, nil
}
