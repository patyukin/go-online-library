package converter

import (
	"github.com/patyukin/go-online-library/internal/repository/model"
	"github.com/patyukin/go-online-library/internal/server/dto"
)

func ToFiltersFromDTO(filters []dto.Filter) []model.Filter {
	var mdls []model.Filter

	for _, filter := range filters {
		m := model.Filter{
			Name:        filter.Name,
			UserName:    filter.UserName,
			BookName:    filter.BookName,
			AuthorName:  filter.AuthorName,
			StartAt:     filter.StartAt,
			NextAfter:   filter.NextAfter,
			PromotionID: filter.PromotionID,
		}

		if filter.ID > 0 {
			m.ID = filter.ID
		}

		mdls = append(mdls, m)
	}

	return mdls
}

func ToFiltersFromDTOWithPromotionID(filters []dto.Filter, promotionID int64) []model.Filter {
	var mdls []model.Filter

	for _, filter := range filters {
		m := model.Filter{
			Name:        filter.Name,
			UserName:    filter.UserName,
			BookName:    filter.BookName,
			AuthorName:  filter.AuthorName,
			StartAt:     filter.StartAt,
			NextAfter:   filter.NextAfter,
			PromotionID: promotionID,
		}

		if filter.ID > 0 {
			m.ID = filter.ID
		}

		mdls = append(mdls, m)
	}

	return mdls
}
