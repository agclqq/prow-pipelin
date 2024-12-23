package repo

import (
	"context"
	"github.com/agclqq/prow-framework/db/repo"
	"github.com/agclqq/prow-pipeline/domain/flow/entity"
	"gorm.io/gorm"
)

type ConfFlow interface {
	Create(ctx context.Context, data any) error
	Delete(ctx context.Context, where any) (int64, error)
	Update(ctx context.Context, where any, data any) (int64, error)
	Pagination(ctx context.Context, columns string, where any, group string, having any, order string, page, pageSize int) (int64, []*entity.ConfFlow, error)
	Select(ctx context.Context, columns string, where any, group string, having any, order string, page, pageSize int) ([]*entity.ConfFlow, error)
	SelectOne(ctx context.Context, columns string, where any, group string, having any, order string) (*entity.ConfFlow, error)
}
type ConfFlowImpl struct {
	db *gorm.DB
}

func NewConfFlowImpl(db *gorm.DB) ConfFlow {
	return &ConfFlowImpl{db: db}
}

func (r *ConfFlowImpl) Table() string {
	return "conf_flows"
}
func (r *ConfFlowImpl) Create(ctx context.Context, data any) error {
	rs := r.db.WithContext(ctx).Table(r.Table()).Create(data)
	return rs.Error
}
func (r *ConfFlowImpl) Delete(ctx context.Context, where any) (int64, error) {
	tx := r.db.WithContext(ctx).Table(r.Table())
	rs := tx.Where(repo.ParseWhere(tx, where)).Delete(&entity.ConfFlow{})
	return rs.RowsAffected, rs.Error
}
func (r *ConfFlowImpl) Update(ctx context.Context, where any, data any) (int64, error) {
	tx := r.db.WithContext(ctx).Table(r.Table())
	tx = repo.ParseWhere(tx, where)
	rs := tx.Updates(data)
	return rs.RowsAffected, rs.Error
}
func (r *ConfFlowImpl) Pagination(ctx context.Context, columns string, where any, group string, having any, order string, page, pageSize int) (int64, []*entity.ConfFlow, error) {
	var data []*entity.ConfFlow
	tx := r.db.WithContext(ctx).Table(r.Table())
	total, tx := repo.Pagination(tx, columns, where, group, having, order, page, pageSize)
	rs := tx.Find(&data)
	return total, data, rs.Error
}
func (r *ConfFlowImpl) Select(ctx context.Context, columns string, where any, group string, having any, order string, page, pageSize int) ([]*entity.ConfFlow, error) {
	var data []*entity.ConfFlow
	tx := r.db.WithContext(ctx).Table(r.Table())
	rs := repo.Select(tx, columns, where, group, having, order, page, pageSize).Find(&data)
	return data, rs.Error
}
func (r *ConfFlowImpl) SelectOne(ctx context.Context, columns string, where any, group string, having any, order string) (*entity.ConfFlow, error) {
	var data entity.ConfFlow
	tx := r.db.WithContext(ctx).Table(r.Table())
	rs := repo.SelectOne(tx, columns, where, group, having, order).Find(&data)
	return &data, rs.Error
}
