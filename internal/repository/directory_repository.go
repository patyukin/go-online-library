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
	tableNameDirectory       = "directories"
	idDirectoryColumn        = "id"
	nameDirectoryColumn      = "name"
	createdAtDirectoryColumn = "created_at"
	updatedAtDirectoryColumn = "updated_at"
	deletedAtDirectoryColumn = "deleted_at"
)

func (r *Repository) InsertDirectory(ctx context.Context, directory model.Directory) (int64, error) {
	directory.CreatedAt = time.Now().UTC()
	builder := sq.Insert(tableNameDirectory).
		PlaceholderFormat(sq.Question).
		Columns(nameDirectoryColumn, createdAtDirectoryColumn).
		Values(directory.Name, directory.CreatedAt)

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, fmt.Errorf("failed to build query: %w", err)
	}

	q := db.Query{Name: "directory_repository.Insert", QueryRaw: query}
	res, err := r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return 0, fmt.Errorf("failed to insert directory: %w", err)
	}

	return res.LastInsertId()
}

func (r *Repository) InsertsDirectories(ctx context.Context, directories []model.Directory) ([]int64, error) {
	ids := make([]int64, 0, len(directories))
	for _, directory := range directories {
		id, err := r.InsertDirectory(ctx, directory)
		if err != nil {
			return nil, fmt.Errorf("failed to insert directory: %w", err)
		}

		ids = append(ids, id)
	}

	return ids, nil
}

func (r *Repository) UpdateDirectory(ctx context.Context, directory model.Directory) error {
	builder := sq.Update(tableNameDirectory).
		PlaceholderFormat(sq.Question).
		Set(nameDirectoryColumn, directory.Name).
		Where(sq.Eq{"id": directory.ID})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{Name: "directory_repository.Update", QueryRaw: query}
	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteDirectory(ctx context.Context, id int64) error {
	builder := sq.Delete(tableNameDirectory).
		PlaceholderFormat(sq.Question).
		Where(sq.Eq{"id": id})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{Name: "directory_repository.Delete", QueryRaw: query}
	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetDirectory(ctx context.Context, id int64) (model.Directory, error) {
	builder := sq.Select(idDirectoryColumn, nameDirectoryColumn, createdAtDirectoryColumn, updatedAtDirectoryColumn).
		From(tableNameDirectory).
		PlaceholderFormat(sq.Question).
		Where(sq.Eq{"id": id})

	query, args, err := builder.ToSql()
	if err != nil {
		return model.Directory{}, err
	}

	q := db.Query{Name: "directory_repository.Get", QueryRaw: query}
	var directory model.Directory
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&directory.ID, &directory.Name, &directory.CreatedAt, &directory.UpdatedAt)
	if err != nil {
		return model.Directory{}, err
	}

	return directory, nil
}

func (r *Repository) GetAllDirectories(ctx context.Context) ([]model.Directory, error) {
	builder := sq.Select(idDirectoryColumn, nameDirectoryColumn, createdAtDirectoryColumn, updatedAtDirectoryColumn).
		From(tableNameDirectory)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{Name: "directory_repository.GetAll", QueryRaw: query}

	rows, err := r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}

	var directories []model.Directory
	for rows.Next() {
		var directory model.Directory
		err = rows.Scan(&directory.ID, &directory.Name, &directory.CreatedAt, &directory.UpdatedAt)
		if err != nil {
			return nil, err
		}
		directories = append(directories, directory)
	}

	return directories, nil
}

func (r *Repository) GetDirectoryByPromotionID(ctx context.Context, promotionID int64) ([]model.Directory, error) {
	builder := sq.Select(idDirectoryColumn, nameDirectoryColumn, createdAtDirectoryColumn, updatedAtDirectoryColumn).
		From(tableNameDirectory).
		Where(sq.Eq{"promotion_id": promotionID})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{Name: "directory_repository.GetByPromotionID", QueryRaw: query}

	rows, err := r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}

	var directories []model.Directory
	for rows.Next() {
		var directory model.Directory
		err = rows.Scan(&directory.ID, &directory.Name, &directory.CreatedAt, &directory.UpdatedAt)
		if err != nil {
			return nil, err
		}

		directories = append(directories, directory)
	}

	return directories, nil
}
