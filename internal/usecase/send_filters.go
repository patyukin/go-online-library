package usecase

import (
	"context"
	"fmt"
	"github.com/patyukin/go-online-library/internal/usecase/model"
)

func (uc *UseCase) SendFilters(ctx context.Context) error {
	var activeFilters []model.Filter
	var err error
	err = uc.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		activeFilters, err = uc.repos.GetActiveFilters(ctx)
		if err != nil {
			return fmt.Errorf("error occured while getting active filters: %w", err)
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("error occured while sending filters: %w", err)
	}

	if len(activeFilters) == 0 {
		return fmt.Errorf("empty list of active filters, nothing to send")
	}

	err = uc.sender.Send(ctx, activeFilters)
	if err != nil {
		return fmt.Errorf("error occured while sending filters: %w", err)
	}

	// recieve response

	return nil
}
