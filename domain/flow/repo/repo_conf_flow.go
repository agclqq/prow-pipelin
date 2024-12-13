package repo

import (
	"context"
	"github.com/agclqq/prow-framework/db/repo"
	"github.com/agclqq/prow-pipeline/domain/flow/agg/ev"
	"gorm.io/gorm"
)

type RepoConfFlow interface {
	Create(ctx context.Context, data any) error
	Delete(ctx context.Context, where any) (int64, error)
	Update(ctx context.Context, where any, data any) (int64, error)
	Pagination(ctx context.Context, columns string, where any, group string, having any, order string, page, pageSize int) (int64, []ev.ConfFlow, error)
	Select(ctx context.Context, columns string, where any, group string, having any, order string, page, pageSize int) ([]ev.ConfFlow, error)
	SelectOne(ctx context.Context, columns string, where any, group string, having any, order string) (*ev.ConfFlow, error)
}
type RepoConfFlowImpl struct {
	db *gorm.DB
}

func NewConfFlow(db *gorm.DB) RepoConfFlow {
	return &RepoConfFlowImpl{db: db}
}

func (r *RepoConfFlowImpl) Table() string {
	return "conf_flows"
}
func (r *RepoConfFlowImpl) Create(ctx context.Context, data any) error {
	rs := r.db.WithContext(ctx).Table(r.Table()).Create(data)
	return rs.Error
}
func (r *RepoConfFlowImpl) Delete(ctx context.Context, where any) (int64, error) {
	tx := r.db.WithContext(ctx).Table(r.Table())
	rs := tx.Where(repo.ParseWhere(tx, where)).Delete(&ev.ConfFlow{})
	return rs.RowsAffected, rs.Error
}
func (r *RepoConfFlowImpl) Update(ctx context.Context, where any, data any) (int64, error) {
	tx := r.db.WithContext(ctx).Table(r.Table())
	tx = repo.ParseWhere(tx, where)
	rs := tx.Updates(data)
	return rs.RowsAffected, rs.Error
}
func (r *RepoConfFlowImpl) Pagination(ctx context.Context, columns string, where any, group string, having any, order string, page, pageSize int) (int64, []ev.ConfFlow, error) {
	var data []ev.ConfFlow
	tx := r.db.WithContext(ctx).Table(r.Table())
	total, tx := repo.Pagination(tx, columns, where, group, having, order, page, pageSize)
	rs := tx.Find(&data)
	return total, data, rs.Error
}
func (r *RepoConfFlowImpl) Select(ctx context.Context, columns string, where any, group string, having any, order string, page, pageSize int) ([]ev.ConfFlow, error) {
	var data []ev.ConfFlow
	tx := r.db.WithContext(ctx).Table(r.Table())
	rs := repo.Select(tx, columns, where, group, having, order, page, pageSize).Find(&data)
	return data, rs.Error
}
func (r *RepoConfFlowImpl) SelectOne(ctx context.Context, columns string, where any, group string, having any, order string) (*ev.ConfFlow, error) {
	var data ev.ConfFlow
	tx := r.db.WithContext(ctx).Table(r.Table())
	rs := repo.SelectOne(tx, columns, where, group, having, order).Find(&data)
	return &data, rs.Error
}
