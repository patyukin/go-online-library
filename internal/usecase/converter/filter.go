package converter

import (
	"github.com/patyukin/go-online-library/internal/handler/reqdto"
	"github.com/patyukin/go-online-library/internal/usecase/model"
	"time"
)

func ToFiltersModelFromReqDTO(dto []reqdto.Filter) []model.Filter {
	var filters []model.Filter
	for _, filter := range dto {
		f := model.Filter{}

		f.MinAge = filter.MinAge
		if filter.MinAge == nil {
			f.MinAge = nil
		}

		f.MaxAge = filter.MaxAge
		if filter.MaxAge == nil {
			f.MaxAge = nil
		}

		f.RegisterDate = filter.RegisterDate
		if filter.RegisterDate.IsZero() {
			f.RegisterDate = nil
		}

		f.LastActivity = filter.LastActivity
		if filter.LastActivity == nil {
			f.LastActivity = nil
		}

		f.NotifyDatetime = filter.NotifyDatetime
		if filter.NotifyDatetime == nil {
			f.NotifyDatetime = nil
		}

		f.CreatedAt = time.Now().UTC()

		filters = append(filters, f)
	}

	return filters
}
