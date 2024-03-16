package filter

import "context"

type FilterRepository interface {
	GetFilters(ctx context.Context) ([]*Filter, error)
}

type Filter struct {
	repo FilterRepository
}

func NewFilter(repo FilterRepository) *Filter {
	return &Filter{
		repo: repo,
	}
}
