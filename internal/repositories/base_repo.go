package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"slices"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type baseRepo[T any] struct {
	bun.IDB
}

func newBaseRepo[T any](IDB bun.IDB) *baseRepo[T] {
	return &baseRepo[T]{
		IDB: IDB,
	}
}

func (in *baseRepo[T]) Insert(ctx context.Context, model *T) error {
	if _, err := in.IDB.NewInsert().Model(model).Exec(ctx); err != nil {
		return in.werr(err)
	}
	return nil
}

func (in *baseRepo[T]) Update(ctx context.Context, model *T) error {
	if _, err := in.IDB.NewUpdate().Model(model).WherePK().Exec(ctx); err != nil {
		return in.werr(err)
	}
	return nil
}

func (in *baseRepo[T]) UpdateColumn(ctx context.Context, model *T, columns ...string) error {
	columns = append(columns, "updated_at")
	if _, err := in.IDB.NewUpdate().Model(model).WherePK().Column(columns...).Exec(ctx); err != nil {
		return in.werr(err)
	}
	return nil
}

func (in *baseRepo[T]) Save(ctx context.Context, model *T) error {
	if _, err := in.IDB.NewInsert().Model(model).On("CONFLICT (id) DO UPDATE").Exec(ctx); err != nil {
		return in.werr(err)
	}
	return nil
}

func (in *baseRepo[T]) FinIDByID(ctx context.Context, id uuid.UUID) (*T, error) {
	model := new(T)
	if err := in.IDB.NewSelect().Model(model).Where("id = ?", id).Scan(ctx, model); err != nil {
		return nil, in.werr(err)
	}
	return model, nil
}

func (in *baseRepo[T]) FinIDBy(ctx context.Context, filters map[string]any) ([]*T, error) {
	return in.SelectMany(ctx, func(q *bun.SelectQuery) *bun.SelectQuery {
		for k, v := range filters {
			q = q.Where(fmt.Sprintf("%s = ?", k), v)
		}
		return q
	})
}

func (in *baseRepo[T]) Exists(ctx context.Context, fn func(q *bun.SelectQuery) *bun.SelectQuery) (bool, error) {
	var model T
	val, err := fn(in.IDB.NewSelect().Model(&model)).Exists(ctx)
	if err != nil {
		return false, in.werr(err)
	}
	return val, nil
}

func (in *baseRepo[T]) Count(ctx context.Context, fn func(q *bun.SelectQuery) *bun.SelectQuery) (int, error) {
	var model T
	val, err := fn(in.IDB.NewSelect().Model(&model)).Count(ctx)
	if err != nil {
		return 0, in.werr(err)
	}
	return val, nil
}

func (in *baseRepo[T]) Delete(ctx context.Context, id uuid.UUID) error {
	var model T
	if _, err := in.IDB.NewDelete().Model(&model).Where("id = ?", id).Exec(ctx); err != nil {
		return in.werr(err)
	}
	return nil
}

func (in *baseRepo[T]) RunInTx(ctx context.Context, opts *sql.TxOptions, f func(ctx context.Context, tx bun.Tx) error) error {
	return in.IDB.RunInTx(ctx, opts, func(ctx context.Context, tx bun.Tx) error {
		return in.werr(f(ctx, tx))
	})
}

type Option int

const OptionUseZeroLenSliceOnNull = iota

type Options []Option

func (in *baseRepo[T]) SelectOne(ctx context.Context, fn func(q *bun.SelectQuery) *bun.SelectQuery) (*T, error) {
	var model T
	err := fn(in.IDB.NewSelect().Model(&model)).Scan(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, sql.ErrNoRows
	} else if err != nil {
		return nil, in.werr(err)
	}

	return &model, nil
}

func (in *baseRepo[T]) SelectMany(ctx context.Context, fn func(q *bun.SelectQuery) *bun.SelectQuery, options ...Option) ([]*T, error) {
	o := Options(options)
	var model []*T
	err := fn(in.IDB.NewSelect().Model(&model)).Scan(ctx)
	if errors.Is(err, sql.ErrNoRows) {

		if slices.Contains(o, OptionUseZeroLenSliceOnNull) {
			model = make([]*T, 0)
			return model, nil
		}

		return nil, sql.ErrNoRows
	} else if err != nil {
		return nil, in.werr(err)
	}

	return model, nil
}

func (in *baseRepo[T]) werr(err error) error {
	return fmt.Errorf("err base repo: %s", err)
}
