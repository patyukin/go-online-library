package converter

import (
	"github.com/patyukin/go-online-library/internal/repository/model"
	"github.com/patyukin/go-online-library/internal/server/dto"
)

func ToDirectoriesFromDTOWithPromotionID(directories []dto.Directory, promotionID int64) []model.Directory {
	var mdls []model.Directory

	for _, directory := range directories {
		m := model.Directory{
			Name:        directory.Name,
			CreatedAt:   directory.CreatedAt,
			UpdatedAt:   directory.UpdatedAt,
			PromotionID: promotionID,
		}

		if directory.ID > 0 {
			m.ID = directory.ID
		}

		mdls = append(mdls, m)
	}

	return mdls
}

func ToDirectoriesFromDTO(directories []dto.Directory) []model.Directory {
	var mdls []model.Directory

	for _, directory := range directories {
		m := model.Directory{
			Name:      directory.Name,
			CreatedAt: directory.CreatedAt,
			UpdatedAt: directory.UpdatedAt,
		}

		if directory.ID > 0 {
			m.ID = directory.ID
		}

		mdls = append(mdls, m)
	}

	return mdls
}
