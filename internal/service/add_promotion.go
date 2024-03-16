package service

import (
	"context"
	"github.com/patyukin/go-online-library/internal/converter"
	"github.com/patyukin/go-online-library/internal/server/dto"
)

type AddPromotion interface {
	AddPromotion(ctx context.Context, promotion dto.Promotion) error
}

func (s *Service) AddPromotion(ctx context.Context, promotion dto.Promotion) error {

	r, err := return s.promotionRepository.CreatePromotion(ctx, converter.ToPromotionFromDTO(&promotion), converter.ToFiltersFromDTO(promotion.Filters), converter.ToDirectoriesFromDTO(promotion.Directories))
	if err != nil {
		return NewNotExistError("Failed to add promotion")
	}

	return nil


}
