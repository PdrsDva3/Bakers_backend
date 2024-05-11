package bread

import (
	"Bakers_backend/internal/entities"
	"Bakers_backend/internal/repository"
	"Bakers_backend/pkg/customerr"
	"context"
	"errors"
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

func (brd BreadRepo) GetBreadByID(ctx context.Context, breadID int) (*entities.Bread, error) {
	var bread entities.Bread
	row := brd.db.QueryRowContext(ctx, `select name, price, description, count, photo from bread where id = $1`, breadID)
	err := row.Scan(&bread.Name, &bread.Price, &bread.Description, &bread.Count, &bread.Photo)
	bread.ID = breadID
	if err != nil {
		return nil, err
	}
	return &bread, nil
}

func (brd BreadRepo) DeleteBread(ctx context.Context, breadID int) error {
	transaction, err := brd.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	result, err := transaction.ExecContext(ctx, `DELETE FROM bread WHERE id=$1;`, breadID)
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

func (brd BreadRepo) ChangeCountBreadByID(ctx context.Context, breadID int, count int64) (int64, error) {
	var oldCount int64
	row := brd.db.QueryRowContext(ctx, `select (count) from bread where id = $1`, breadID)
	err := row.Scan(&oldCount)
	if err != nil {
		return 0, err
	}
	transaction, err := brd.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}
	if oldCount+count < 0 {
		return 0, customerr.ErrorMessage(2, "out of count bread")
	}
	result, err := transaction.ExecContext(ctx, `UPDATE bread SET count = $2 WHERE id = $1;`, breadID, count+oldCount)
	if err != nil {
		return 0, err
	}
	cnt, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	if cnt != 1 {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return 0, errors.New("failed rollback")
		}
		return 0, errors.New("failed to update password")
	}
	if err = transaction.Commit(); err != nil {
		return 0, err
	}
	return count + oldCount, nil
}
