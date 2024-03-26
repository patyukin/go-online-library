package usecase

import (
	"context"
	"github.com/patyukin/go-online-library/internal/usecase/model"
	"github.com/patyukin/go-online-library/pkg/db"
)

//go:generate go run github.com/vektra/mockery/v2@v2.42.0 --name=Repositories --with-expecter=true
//go:generate go run github.com/vektra/mockery/v2@v2.42.0 --name=TxManager --with-expecter=true

type Repositories interface {
	InsertPromotion(ctx context.Context, promotion model.Promotion) (int64, error)
	UpdatePromotion(ctx context.Context, promotion model.Promotion) error
	DeletePromotion(ctx context.Context, filterID int64) error
	GetPromotion(ctx context.Context, filterID int64) (model.Promotion, error)
	GetAllPromotions(ctx context.Context) ([]model.Promotion, error)
	InsertsPromotionDirectories(ctx context.Context, promotionID int64, directoryIDs []int64) error
	InsertsPromotionFilters(ctx context.Context, promotionID int64, filterIDs []int64) error

	InsertFilter(ctx context.Context, filter model.Filter) (int64, error)
	InsertsFilters(ctx context.Context, filters []model.Filter) ([]int64, error)
	UpdateFilter(ctx context.Context, filter model.Filter) error
	DeleteFilter(ctx context.Context, filterID int64) error
	GetFilter(ctx context.Context, filterID int64) (model.Filter, error)
	GetAllFilters(ctx context.Context) ([]model.Filter, error)

	InsertDirectory(ctx context.Context, directory model.Directory) (int64, error)
	InsertsDirectories(ctx context.Context, directories []model.Directory) ([]int64, error)
	UpdateDirectory(ctx context.Context, directory model.Directory) error
	DeleteDirectory(ctx context.Context, id int64) error
	GetDirectory(ctx context.Context, id int64) (model.Directory, error)
	GetAllDirectories(ctx context.Context) ([]model.Directory, error)
	GetDirectoryByPromotionID(ctx context.Context, promotionID int64) ([]model.Directory, error)

	GetActiveFilters(ctx context.Context) ([]model.Filter, error)
}

type Cacher interface {
	GetPromotion(ctx context.Context, id int64) (model.Promotion, error)
}

type TxManager interface {
	ReadCommitted(ctx context.Context, f db.Handler) error
}

type Sender interface {
	Send(ctx context.Context, filters []model.Filter) error
}

type UseCase struct {
	repos     Repositories
	txManager TxManager
	sender    Sender
}

func (uc *UseCase) GetPromotion(ctx context.Context, id int64) (model.Promotion, error) {
	//TODO implement me
	panic("implement me")
}

func New(repos Repositories, txManager TxManager, sender Sender) *UseCase {
	return &UseCase{
		repos:     repos,
		txManager: txManager,
		sender:    sender,
	}
}
