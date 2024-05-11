package breadserv

import (
	"Bakers_backend/internal/entities"
	"Bakers_backend/internal/repository"
	"Bakers_backend/internal/service"
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
	return bread, nil
}
