package repository

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/patyukin/go-online-library/internal/repository/model"
	"github.com/patyukin/go-online-library/pkg/db"
)

const (
	tableNamePromo       = "promotions"
	idPromoColumn        = "id"
	namePromoColumn      = "name"
	activePromoColumn    = "active"
	countPromoColumn     = "count"
	createdAtPromoColumn = "created_at"
	updatedAtPromoColumn = "updated_at"
	deletedAtPromoColumn = "deleted_at"
)

type PromotionRepo struct {
	db db.Client
}

func NewPromotionRepo(db db.Client) *PromotionRepo {
	return &PromotionRepo{
		db: db,
	}
}

func (r *PromotionRepo) Insert(ctx context.Context, promotion model.Promotion) (int64, error) {
	builder := sq.Insert(tableNamePromo).
		Columns(namePromoColumn, activePromoColumn, countPromoColumn).
		Values(promotion.Name, promotion.Active, promotion.Count).
		PlaceholderFormat(sq.Question).
		Suffix("RETURNING " + idPromoColumn)

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{Name: "promotion_repository.Insert", QueryRaw: query}
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&promotion.ID)
	if err != nil {
		return 0, err
	}

	return promotion.ID, nil
}

func (r *PromotionRepo) Update(ctx context.Context, promotion model.Promotion) error {
	builder := sq.Update(tableNamePromo).
		PlaceholderFormat(sq.Question).
		Set(namePromoColumn, promotion.Name).
		Set("active", promotion.Active).
		Set("count", promotion.Count).
		Where(sq.Eq{"id": promotion.ID})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{Name: "promotion_repository.Update", QueryRaw: query}
	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *PromotionRepo) Delete(ctx context.Context, filterID int64) error {
	builder := sq.Delete(tableNamePromo).
		PlaceholderFormat(sq.Question).
		Where(sq.Eq{"id": filterID})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{Name: "promotion_repository.Delete", QueryRaw: query}
	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *PromotionRepo) Get(ctx context.Context, filterID int64) (model.Promotion, error) {
	builder := sq.Select(idPromoColumn, namePromoColumn, activePromoColumn, countPromoColumn, createdAtPromoColumn, updatedAtPromoColumn, deletedAtPromoColumn).
		From(tableNamePromo).
		PlaceholderFormat(sq.Question).
		Where(sq.Eq{"id": filterID})

	query, args, err := builder.ToSql()
	if err != nil {
		return model.Promotion{}, err
	}

	q := db.Query{Name: "promotion_repository.Get", QueryRaw: query}

	var promotion model.Promotion
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&promotion.ID, &promotion.Name, &promotion.Active, &promotion.Count, &promotion.CreatedAt, &promotion.UpdatedAt, &promotion.DeletedAt)
	if err != nil {
		return model.Promotion{}, err
	}

	return promotion, nil
}

func (r *PromotionRepo) GetAll(ctx context.Context) ([]model.Promotion, error) {
	builder := sq.Select(idPromoColumn, namePromoColumn, activePromoColumn, countPromoColumn, createdAtPromoColumn, updatedAtPromoColumn, deletedAtPromoColumn).
		From(tableNamePromo)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{Name: "promotion_repository.GetAll", QueryRaw: query}

	rows, err := r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}

	var promotions []model.Promotion
	for rows.Next() {
		var promotion model.Promotion
		err = rows.Scan(&promotion.ID, &promotion.Name, &promotion.Active, &promotion.Count, &promotion.CreatedAt, &promotion.UpdatedAt, &promotion.DeletedAt)
		if err != nil {
			return nil, err
		}

		promotions = append(promotions, promotion)
	}

	return promotions, nil
}
