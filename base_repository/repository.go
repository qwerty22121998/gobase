package base_repository

import (
	"context"
	"github.com/qwerty22121998/gobase/base_model"
	"github.com/qwerty22121998/gobase/base_util"
	"github.com/qwerty22121998/gobase/pagination"
	"github.com/qwerty22121998/gobase/preload"
	"github.com/qwerty22121998/gobase/query"
	"gorm.io/gorm"
)

const (
	ContextDBKey = "_tx"
)

type IRepository[T base_model.IModel] interface {
}

type repository[T base_model.IModel] struct {
	db *gorm.DB
}

func New[T base_model.IModel](db *gorm.DB) IRepository[T] {
	return &repository[T]{db: db}
}

func (r *repository[T]) DB(ctx context.Context) *gorm.DB {
	db, ok := ctx.Value(ContextDBKey).(*gorm.DB)
	if !ok {
		return r.db
	}
	return db
}

func (r *repository[T]) Create(ctx context.Context, data T) error {
	data.SetCreated(base_util.GetUser(ctx), base_util.GetNow(ctx))
	data.SetUpdated(base_util.GetUser(ctx), base_util.GetNow(ctx))
	return r.DB(ctx).Create(data).Error
}

func (r *repository[T]) Save(ctx context.Context, data T) error {
	data.SetUpdated(base_util.GetUser(ctx), base_util.GetNow(ctx))
	return r.DB(ctx).Save(data).Error
}

func (r *repository[T]) Delete(ctx context.Context, data T) error {
	return r.DB(ctx).Delete(data).Error
}

func (r *repository[T]) FindByID(ctx context.Context, id uint) (T, error) {
	var result T
	err := r.DB(ctx).First(&result, id).Error
	return result, err
}

func (r *repository[T]) FindMany(ctx context.Context, q query.Condition, p *pagination.Pagination, preloads ...preload.Opt) ([]T, error) {
	res := make([]T, 0)
	zero := new(T)
	db := r.DB(ctx).Model(zero)
	db = db.Scopes(q)
	db = preload.Group(preloads...).Apply(db)
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, err
	}

	if err := p.Apply(db).Find(&res).Error; err != nil {
		return nil, err
	}
	p.SetTotal(total)
	return res, nil
}
