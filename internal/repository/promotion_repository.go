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
	tableNamePromo         = "promotions"
	idPromoColumn          = "id"
	namePromoColumn        = "name"
	descriptionPromoColumn = "description"
	commentPromoColumn     = "comment"
	statusPromoColumn      = "status"
	typePromoColumn        = "type"
	createdAtPromoColumn   = "created_at"
	updatedAtPromoColumn   = "updated_at"
	deletedAtPromoColumn   = "deleted_at"
)

func (r *Repository) InsertPromotion(ctx context.Context, promotion model.Promotion) (int64, error) {
	promotion.CreatedAt = time.Now().UTC()
	qb := sq.Insert(tableNamePromo).
		PlaceholderFormat(sq.Question).
		Columns(namePromoColumn, descriptionPromoColumn, commentPromoColumn, statusPromoColumn, typePromoColumn, createdAtPromoColumn, updatedAtPromoColumn, deletedAtPromoColumn).
		Values(promotion.Name, promotion.Description, promotion.Comment, promotion.Status, promotion.Type, promotion.CreatedAt, nil, nil)

	query, args, err := qb.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{Name: "promotion_repository.InsertPromotion", QueryRaw: query}

	res, err := r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return 0, fmt.Errorf("promotion_repository.InsertPromotion, r.db.DB().ExecContext: %w", err)
	}

	return res.LastInsertId()
}

func (r *Repository) UpdatePromotion(ctx context.Context, promotion model.Promotion) error {

	return nil
}

func (r *Repository) DeletePromotion(ctx context.Context, filterID int64) error {
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

func (r *Repository) GetPromotion(ctx context.Context, filterID int64) (model.Promotion, error) {
	return model.Promotion{}, nil
}

func (r *Repository) GetAllPromotions(ctx context.Context) ([]model.Promotion, error) {
	return []model.Promotion{}, nil
}

func (r *Repository) InsertsPromotionDirectories(ctx context.Context, promotionID int64, directoryIDs []int64) error {
	qb := sq.Insert("promotions_directories").
		PlaceholderFormat(sq.Question).
		Columns("promotion_id", "directory_id")

	for _, d := range directoryIDs {
		qb = qb.Values(promotionID, d)
	}

	query, args, err := qb.ToSql()
	if err != nil {
		return fmt.Errorf("promotion_repository.InsertsPromotionDirectories, qb.ToSql: %w", err)
	}

	q := db.Query{Name: "promotion_repository.InsertsPromotionDirectories", QueryRaw: query}
	res, err := r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return fmt.Errorf("promotion_repository.InsertsPromotionDirectories, r.db.DB().ExecContext: %w", err)
	}

	count, err := res.RowsAffected()
	if err != nil || count != int64(len(directoryIDs)) {
		return fmt.Errorf("promotion_repository.InsertsPromotionDirectories, res.RowsAffected: %w", err)
	}

	return nil
}

func (r *Repository) InsertsPromotionFilters(ctx context.Context, promotionID int64, filterIDs []int64) error {
	qb := sq.Insert("promotions_filters").
		PlaceholderFormat(sq.Question).
		Columns("promotion_id", "filter_id")

	for _, f := range filterIDs {
		qb = qb.Values(promotionID, f)
	}

	query, args, err := qb.ToSql()
	if err != nil {
		return fmt.Errorf("promotion_repository.InsertsPromotionFilters, qb.ToSql: %w", err)
	}

	q := db.Query{Name: "promotion_repository.InsertsPromotionFilters", QueryRaw: query}
	res, err := r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return fmt.Errorf("promotion_repository.InsertsPromotionFilters, r.db.DB().ExecContext: %w", err)
	}

	count, err := res.RowsAffected()
	if err != nil || count != int64(len(filterIDs)) {
		return fmt.Errorf("promotion_repository.InsertsPromotionFilters, res.RowsAffected: %w", err)
	}

	return nil
}
