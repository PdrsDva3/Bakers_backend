package breadserv

import (
	"Bakers_backend/internal/entities"
	"Bakers_backend/internal/repository"
	"Bakers_backend/internal/service"
	cerr "Bakers_backend/pkg/customerr"
	"context"
)

func InitBreadService(breadRepo repository.BreadRepo) service.BreadService {
	return BreadService{BreadRepo: breadRepo}
}

type BreadService struct {
	BreadRepo repository.BreadRepo
}

func (brd BreadService) BreadCreate(ctx context.Context, breadCreate entities.BreadBase) (int, error) {
	id, err := brd.BreadRepo.Create(ctx, breadCreate)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (brd BreadService) GetBread(ctx context.Context, breadID int) (*entities.Bread, error) {
	bread, err := brd.BreadRepo.GetBreadByID(ctx, breadID)
	if err != nil {
		return nil, err
	}
	if bread == nil {
		return nil, cerr.Err(cerr.Bread, cerr.Service, cerr.NotFound, nil).Error()
	}

	return bread, nil
}

func (brd BreadService) ChangeBread(ctx context.Context, breadID int, count int) (int, error) {
	newCount, err := brd.BreadRepo.ChangeCountBreadByID(ctx, breadID, count)
	if err != nil {
		return 0, err
	}
	return newCount, nil
}

func (brd BreadService) DeleteBread(ctx context.Context, breadID int) error {
	err := brd.BreadRepo.DeleteBread(ctx, breadID)
	if err != nil {
		return err
	}
	return nil
}
