package converter

import (
	"github.com/patyukin/go-online-library/internal/handler/reqdto"
	"github.com/patyukin/go-online-library/internal/usecase/model"
	"gopkg.in/guregu/null.v3"
)

func ToPromotionModelFromReqDTO(dto reqdto.Promotion) model.Promotion {
	m := model.Promotion{
		Name:        dto.Name,
		Description: dto.Description,
		Status:      dto.Status,
		Type:        dto.Type,
	}

	m.Comment = null.StringFromPtr(dto.Comment)

	if dto.ID != 0 {
		m.ID = dto.ID
	}

	return m
}
