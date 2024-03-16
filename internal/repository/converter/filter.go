package converter

import (
	"github.com/patyukin/go-online-library/internal/repository/model"
	repo "github.com/patyukin/go-online-library/internal/repository/model"
)

func ToFilterFromRepo(filter *repo.Filter) *model.Filter {
	return &model.Filter{
		ID:        filter.ID,
		Name:      filter.Name,
		CreatedAt: filter.CreatedAt,
		UpdatedAt: filter.UpdatedAt,
		DeletedAt: filter.DeletedAt,
	}
}
