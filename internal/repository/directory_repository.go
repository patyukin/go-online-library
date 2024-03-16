package repository

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/patyukin/go-online-library/internal/repository/model"
	"github.com/patyukin/go-online-library/pkg/db"
)

const (
	tableNameDirectory         = "directories"
	idDirectoryColumn          = "id"
	nameDirectoryColumn        = "name"
	promotionIDDirectoryColumn = "promotion_id"
	createdAtDirectoryColumn   = "created_at"
	updatedAtDirectoryColumn   = "updated_at"
)

type DirectoryRepo struct {
	db db.Client
}

func NewDirectoryRepo(db db.Client) *DirectoryRepo {
	return &DirectoryRepo{
		db: db,
	}
}

func (r *DirectoryRepo) Insert(ctx context.Context, directory model.Directory) (int64, error) {
	builder := sq.Insert(tableNameDirectory).
		Columns(nameDirectoryColumn, promotionIDDirectoryColumn).
		Values(directory.Name, directory.PromotionID).
		PlaceholderFormat(sq.Question).
		Suffix("RETURNING " + idDirectoryColumn)

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{Name: "directory_repository.Insert", QueryRaw: query}
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&directory.ID)
	if err != nil {
		return 0, err
	}

	return directory.ID, nil
}

func (r *DirectoryRepo) Inserts(ctx context.Context, directories []model.Directory) (int64, error) {
	builder := sq.Insert(tableNameDirectory).
		Columns(nameDirectoryColumn, promotionIDDirectoryColumn).
		PlaceholderFormat(sq.Question).
		Suffix("RETURNING " + idDirectoryColumn)

	for _, directory := range directories {
		builder = builder.Values(directory.Name, directory.PromotionID)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{Name: "directory_repository.Inserts", QueryRaw: query}
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&directories[0].ID)
	if err != nil {
		return 0, err
	}

	return int64(len(directories)), nil
}

func (r *DirectoryRepo) Update(ctx context.Context, directory model.Directory) error {
	builder := sq.Update(tableNameDirectory).
		PlaceholderFormat(sq.Question).
		Set(nameDirectoryColumn, directory.Name).
		Set(promotionIDDirectoryColumn, directory.PromotionID).
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

func (r *DirectoryRepo) Delete(ctx context.Context, id int64) error {
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

func (r *DirectoryRepo) Get(ctx context.Context, id int64) (model.Directory, error) {
	builder := sq.Select(idDirectoryColumn, nameDirectoryColumn, promotionIDDirectoryColumn, createdAtDirectoryColumn, updatedAtDirectoryColumn).
		From(tableNameDirectory).
		PlaceholderFormat(sq.Question).
		Where(sq.Eq{"id": id})

	query, args, err := builder.ToSql()
	if err != nil {
		return model.Directory{}, err
	}

	q := db.Query{Name: "directory_repository.Get", QueryRaw: query}
	var directory model.Directory
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&directory.ID, &directory.Name, &directory.PromotionID, &directory.CreatedAt, &directory.UpdatedAt)
	if err != nil {
		return model.Directory{}, err
	}

	return directory, nil
}

func (r *DirectoryRepo) GetAll(ctx context.Context) ([]model.Directory, error) {
	builder := sq.Select(idDirectoryColumn, nameDirectoryColumn, promotionIDDirectoryColumn, createdAtDirectoryColumn, updatedAtDirectoryColumn).
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
		err = rows.Scan(&directory.ID, &directory.Name, &directory.PromotionID, &directory.CreatedAt, &directory.UpdatedAt)
		if err != nil {
			return nil, err
		}
		directories = append(directories, directory)
	}

	return directories, nil
}

func (r *DirectoryRepo) GetByPromotionID(ctx context.Context, promotionID int64) ([]model.Directory, error) {
	builder := sq.Select(idDirectoryColumn, nameDirectoryColumn, promotionIDDirectoryColumn, createdAtDirectoryColumn, updatedAtDirectoryColumn).
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
		err = rows.Scan(&directory.ID, &directory.Name, &directory.PromotionID, &directory.CreatedAt, &directory.UpdatedAt)
		if err != nil {
			return nil, err
		}

		directories = append(directories, directory)
	}

	return directories, nil
}
