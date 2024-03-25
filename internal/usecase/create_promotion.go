package usecase

import (
	"context"
	"fmt"
	"github.com/patyukin/go-online-library/internal/usecase/model"
)

func (uc *UseCase) CreatePromotion(ctx context.Context, promotion model.Promotion, filters []model.Filter, directories []model.Directory) error {
	err := uc.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		id, err := uc.repos.InsertPromotion(ctx, promotion)
		if err != nil {
			return err
		}

		ids, err := uc.repos.InsertsFilters(ctx, filters)
		if err != nil {
			return fmt.Errorf("failed to insert filters: %w", err)
		}

		err = uc.repos.InsertsPromotionFilters(ctx, id, ids)
		if err != nil {
			return fmt.Errorf("failed to insert promotion directories: %w", err)
		}

		ids, err = uc.repos.InsertsDirectories(ctx, directories)
		if err != nil {
			return err
		}

		err = uc.repos.InsertsPromotionDirectories(ctx, id, ids)
		if err != nil {
			return fmt.Errorf("failed to insert promotion directories: %w", err)
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
