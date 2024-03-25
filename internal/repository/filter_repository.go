package repository

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/patyukin/go-online-library/internal/usecase/model"
	"github.com/patyukin/go-online-library/pkg/db"
	"time"
)

const (
	tableFilterName            = "filters"
	idFilterColumn             = "id"
	minAgeFilterColumn         = "min_age"
	maxAgeFilterColumn         = "max_age"
	registerDateFilterColumn   = "register_date"
	lastActivityFilterColumn   = "last_activity"
	notifyDatetimeFilterColumn = "notify_datetime"
	createdAtFilterColumn      = "created_at"
	updatedAtFilterColumn      = "updated_at"
	deletedAtFilterColumn      = "deleted_at"
)

func (r *Repository) InsertFilter(ctx context.Context, filter model.Filter) (int64, error) {
	filter.CreatedAt = time.Now().UTC()
	builder := sq.Insert(tableFilterName).
		PlaceholderFormat(sq.Question).
		Columns(minAgeFilterColumn, maxAgeFilterColumn, registerDateFilterColumn, lastActivityFilterColumn, notifyDatetimeFilterColumn, createdAtFilterColumn).
		Values(filter.MinAge, filter.MaxAge, filter.RegisterDate, filter.LastActivity, filter.NotifyDatetime, filter.CreatedAt)

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, fmt.Errorf("error occured while inserting filter: %w", err)
	}

	q := db.Query{Name: "filter_repository.Insert", QueryRaw: query}
	res, err := r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return 0, fmt.Errorf("error occured while inserting filter: %w", err)
	}

	return res.LastInsertId()
}

func (r *Repository) InsertsFilters(ctx context.Context, filters []model.Filter) ([]int64, error) {
	ids := make([]int64, 0, len(filters))
	for _, filter := range filters {
		id, err := r.InsertFilter(ctx, filter)
		if err != nil {
			return nil, fmt.Errorf("error occured while inserting filter: %v", err)
		}

		ids = append(ids, id)
	}

	return ids, nil
}

func (r *Repository) UpdateFilter(ctx context.Context, filter model.Filter) error {

	return nil
}

func (r *Repository) DeleteFilter(ctx context.Context, filterID int64) error {
	builder := sq.Delete(tableFilterName).
		PlaceholderFormat(sq.Question).
		Where(sq.Eq{"id": filterID})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{Name: "filter_repository.Delete", QueryRaw: query}
	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetFilter(ctx context.Context, filterID int64) (model.Filter, error) {
	return model.Filter{}, nil
}

func (r *Repository) GetAllFilters(ctx context.Context) ([]model.Filter, error) {
	return []model.Filter{}, nil
}

func (r *Repository) GetActiveFilters(ctx context.Context) ([]model.Filter, error) {
	builder := sq.Select(idFilterColumn).
		From(tableFilterName).
		Where(sq.Eq{deletedAtFilterColumn: nil})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("error occured while getting active filters: %w", err)
	}

	q := db.Query{Name: "filter_repository.GetActiveFilters", QueryRaw: query}
	rows, err := r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return nil, fmt.Errorf("error occured while getting active filters: %w", err)
	}

	var filters []model.Filter

	for rows.Next() {
		var filter model.Filter
		if err = rows.Scan(&filter); err != nil {
			return nil, fmt.Errorf("error occured while getting active filters: %w", err)
		}

		filters = append(filters, filter)
	}

	return filters, nil
}
