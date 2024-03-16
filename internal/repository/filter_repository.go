package repository

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/patyukin/go-online-library/internal/repository/model"
	"github.com/patyukin/go-online-library/pkg/db"
)

const (
	tableFilterName         = "filters"
	idFilterColumn          = "id"
	nameFilterColumn        = "name"
	userNameFilterColumn    = "user_name"
	bookNameFilterColumn    = "book_name"
	authorNameFilterColumn  = "author_name"
	startAtFilterColumn     = "start_at"
	nextAfterFilterColumn   = "next_after"
	promotionIDFilterColumn = "promotion_id"
	createdAtFilterColumn   = "created_at"
	updatedAtFilterColumn   = "updated_at"
	deletedAtFilterColumn   = "deleted_at"
)

type FilterRepo struct {
	db db.Client
}

func NewFilterRepo(db db.Client) *FilterRepo {
	return &FilterRepo{
		db: db,
	}
}

func (r *FilterRepo) Insert(ctx context.Context, filter model.Filter) (int64, error) {
	builder := sq.Insert(tableFilterName).
		PlaceholderFormat(sq.Question).
		Columns(nameFilterColumn, userNameFilterColumn, bookNameFilterColumn, authorNameFilterColumn, startAtFilterColumn, nextAfterFilterColumn, promotionIDFilterColumn).
		Values(filter.Name, filter.UserName, filter.BookName, filter.AuthorName, filter.StartAt, filter.NextAfter, filter.PromotionID).
		Suffix("RETURNING " + idFilterColumn)

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{Name: "filter_repository.Insert", QueryRaw: query}
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&filter.ID)
	if err != nil {
		return 0, err
	}

	return filter.ID, nil
}

func (r *FilterRepo) Inserts(ctx context.Context, filters []model.Filter) (int64, error) {
	builder := sq.Insert(tableFilterName).
		PlaceholderFormat(sq.Question).
		Columns(nameFilterColumn, userNameFilterColumn, bookNameFilterColumn, authorNameFilterColumn, startAtFilterColumn, nextAfterFilterColumn, promotionIDFilterColumn).
		Values(filters[0].Name, filters[0].UserName, filters[0].BookName, filters[0].AuthorName, filters[0].StartAt, filters[0].NextAfter, filters[0].PromotionID)

	for i := 1; i < len(filters); i++ {
		builder = builder.Values(filters[i].Name, filters[i].UserName, filters[i].BookName, filters[i].AuthorName, filters[i].StartAt, filters[i].NextAfter, filters[i].PromotionID)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{Name: "filter_repository.Inserts", QueryRaw: query}
	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return 0, err
	}

	return int64(len(filters)), nil
}

func (r *FilterRepo) Update(ctx context.Context, filter model.Filter) error {
	builder := sq.Update(tableFilterName).
		PlaceholderFormat(sq.Question).
		Set(nameFilterColumn, filter.Name).
		Set(userNameFilterColumn, filter.UserName).
		Set(bookNameFilterColumn, filter.BookName).
		Set(authorNameFilterColumn, filter.AuthorName).
		Set(startAtFilterColumn, filter.StartAt).
		Set(nextAfterFilterColumn, filter.NextAfter).
		Set(promotionIDFilterColumn, filter.PromotionID).
		Where(sq.Eq{"id": filter.ID})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{Name: "filter_repository.Update", QueryRaw: query}
	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *FilterRepo) Delete(ctx context.Context, filterID int64) error {
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

func (r *FilterRepo) Get(ctx context.Context, filterID int64) (model.Filter, error) {
	builder := sq.Select(idFilterColumn, nameFilterColumn, userNameFilterColumn, bookNameFilterColumn, authorNameFilterColumn, startAtFilterColumn, nextAfterFilterColumn, promotionIDFilterColumn, deletedAtFilterColumn).
		From(tableFilterName).
		PlaceholderFormat(sq.Question).
		Where(sq.Eq{"id": filterID})

	query, args, err := builder.ToSql()
	if err != nil {
		return model.Filter{}, err
	}

	q := db.Query{Name: "filter_repository.Get", QueryRaw: query}
	var filter model.Filter
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&filter.ID, &filter.Name, &filter.UserName, &filter.BookName, &filter.AuthorName, &filter.StartAt, &filter.NextAfter, &filter.PromotionID, &filter.DeletedAt)
	if err != nil {
		return model.Filter{}, err
	}

	return filter, nil
}

func (r *FilterRepo) GetAll(ctx context.Context) ([]model.Filter, error) {
	builder := sq.Select(idFilterColumn, nameFilterColumn, userNameFilterColumn, bookNameFilterColumn, authorNameFilterColumn, startAtFilterColumn, nextAfterFilterColumn, promotionIDFilterColumn, createdAtFilterColumn, updatedAtFilterColumn, deletedAtFilterColumn).
		From(tableFilterName).
		PlaceholderFormat(sq.Question)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{Name: "filter_repository.GetAll", QueryRaw: query}
	rows, err := r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			//TODO
		}
	}(rows)

	var filters []model.Filter
	for rows.Next() {
		var filter model.Filter
		err = rows.Scan(&filter.ID, &filter.Name, &filter.UserName, &filter.BookName, &filter.AuthorName, &filter.StartAt, &filter.NextAfter, &filter.PromotionID, &filter.CreatedAt, &filter.UpdatedAt, &filter.DeletedAt)
		if err != nil {
			return nil, err
		}

		filters = append(filters, filter)
	}

	return filters, nil
}
