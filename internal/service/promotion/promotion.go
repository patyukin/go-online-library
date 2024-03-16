package promotion

import (
	"context"
	"github.com/patyukin/go-online-library/internal/repository/model"
	"github.com/patyukin/go-online-library/pkg/db"
)

type PromotionRepo interface {
	Insert(ctx context.Context, promotion model.Promotion) (int64, error)
	Update(ctx context.Context, promotion model.Promotion) error
	Delete(ctx context.Context, filterID int64) error
	Get(ctx context.Context, filterID int64) (model.Promotion, error)
	GetAll(ctx context.Context) ([]model.Promotion, error)
}

type FilterRepo interface {
	Insert(ctx context.Context, filter model.Filter) (int64, error)
	Inserts(ctx context.Context, filters []model.Filter) (int64, error)
	Update(ctx context.Context, filter model.Filter) error
	Delete(ctx context.Context, filterID int64) error
	Get(ctx context.Context, filterID int64) (model.Filter, error)
	GetAll(ctx context.Context) ([]model.Filter, error)
}

type DirectoryRepo interface {
	Insert(ctx context.Context, directory model.Directory) (int64, error)
	Inserts(ctx context.Context, directories []model.Directory) (int64, error)
	Update(ctx context.Context, directory model.Directory) error
	Delete(ctx context.Context, id int64) error
	Get(ctx context.Context, id int64) (model.Directory, error)
	GetAll(ctx context.Context) ([]model.Directory, error)
	GetByPromotionID(ctx context.Context, promotionID int64) ([]model.Directory, error)
}

type PromotionService struct {
	promotionRepo PromotionRepo
	filterRepo    FilterRepo
	directoryRepo DirectoryRepo
	txManager     db.TxManager
}

func New(promotionRepo PromotionRepo, filterRepo FilterRepo, directoryRepo DirectoryRepo, txManager db.TxManager) *PromotionService {
	return &PromotionService{
		filterRepo:    filterRepo,
		promotionRepo: promotionRepo,
		directoryRepo: directoryRepo,
		txManager:     txManager,
	}
}

func (s *PromotionService) CreatePromotion(ctx context.Context, promotion model.Promotion, filters []model.Filter, directories []model.Directory) error {
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		_, err := s.promotionRepo.Insert(ctx, promotion)
		if err != nil {
			return err
		}

		_, err = s.filterRepo.Inserts(ctx, filters)
		if err != nil {
			return err
		}

		_, err = s.directoryRepo.Inserts(ctx, directories)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *PromotionService) GetAllPromotions(ctx context.Context) error {

	return nil
}
