package converter

import (
	"github.com/patyukin/go-online-library/internal/repository/model"
	"github.com/patyukin/go-online-library/internal/server/dto"
)

func ToPromotionFromDTO(promotion *dto.Promotion) model.Promotion {
	r := model.Promotion{
		Name:   promotion.Name,
		Active: promotion.Active,
		Count:  promotion.Count,
	}

	if promotion.ID > 0 {
		r.ID = promotion.ID
	}

	return r
}
