package repo

import (
	"context"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/ozoncp/ocp-instruction-api/internal/models"
	"github.com/ozoncp/ocp-instruction-api/pkg/db"
)

// Repo - storage interface for models.Instruction
type Repo interface {
	Add(ctx context.Context, entities []models.Instruction) error
	List(ctx context.Context, limit, offset uint64) ([]models.Instruction, error)
	Describe(ctx context.Context, id uint64) (*models.Instruction, error)
	Remove(ctx context.Context, id uint64) error
}

func NewRepo() Repo {
	return &repo{}
}

const (
	tablename = "instruction"
)

var (
	ErrNotFound = errors.New("Not found")
)

type repo struct{}

func (r *repo) Add(ctx context.Context, entities []models.Instruction) error {
	query := sq.Insert(tablename).
		Columns("instruction_id", "classroom_id", "text", "prev_id").
		PlaceholderFormat(sq.Dollar).
		RunWith(db.GetDB(ctx))

	for _, ent := range entities {
		query = query.Values(ent.Id, ent.ClassroomId, ent.Text, ent.PrevId)
	}

	_, err := query.ExecContext(ctx)

	return err
}

func (r *repo) List(ctx context.Context, limit, offset uint64) ([]models.Instruction, error) {
	query := sq.Select("instruction_id", "text", "prev_id", "classroom_id").
		From(tablename).
		OrderBy("id").
		RunWith(db.GetDB(ctx)).
		PlaceholderFormat(sq.Dollar)

	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}

	rows, err := query.QueryContext(ctx)

	if err != nil {
		return nil, err
	}

	var ents []models.Instruction
	for rows.Next() {
		var ent models.Instruction
		if err := rows.Scan(
			&ent.Id,
			&ent.Text,
			&ent.PrevId,
			&ent.ClassroomId,
		); err != nil {
			continue
		}
		ents = append(ents, ent)
	}

	return ents, nil
}

func (r *repo) Describe(ctx context.Context, id uint64) (*models.Instruction, error) {
	query := sq.Select("instruction_id", "text", "prev_id", "classroom_id").
		From(tablename).
		Where(sq.Eq{"instruction_id": id}).
		RunWith(db.GetDB(ctx)).
		PlaceholderFormat(sq.Dollar)

	var ret models.Instruction
	if err := query.QueryRowContext(ctx).Scan(&ret.Id, &ret.Text, &ret.PrevId, &ret.ClassroomId); err != nil {
		return nil, err
	}
	return &ret, nil
}

func (r *repo) Remove(ctx context.Context, id uint64) error {
	res, err := sq.Delete(tablename).
		Where(sq.Eq{"instruction_id": id}).
		RunWith(db.GetDB(ctx)).
		PlaceholderFormat(sq.Dollar).
		ExecContext(ctx)

	if err != nil {
		return err
	}

	cnt, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if cnt < 1 {
		return ErrNotFound
	}

	return nil
}
