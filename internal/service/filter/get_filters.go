package filter

import "context"

func (s *Filter) GetFilters(ctx context.Context) ([]*Filter, error) {
	return s.repo.GetFilters(ctx)
}
