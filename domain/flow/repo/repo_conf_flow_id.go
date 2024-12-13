package repo

import (
	"context"
	"github.com/agclqq/prow-framework/db/repo"
	"github.com/agclqq/prow-pipeline/domain/flow/agg/ev"
	"gorm.io/gorm"
)

type RepoConfFlowId interface {
	Create(ctx context.Context, data any) error
	Delete(ctx context.Context, where any) (int64, error)
	Update(ctx context.Context, where *any, data any) (int64, error)
	SelectOne(columns string, where any, group string, having any, order string) (*ev.ConfFlowId, error)
}
type RepoConfFlowIdImpl struct {
	db *gorm.DB
}

func NewRepoConfFlowIdImpl(db *gorm.DB) RepoConfFlowId {
	return &RepoConfFlowIdImpl{db: db}
}

func (r *RepoConfFlowIdImpl) Table() string {
	return "conf_flow_ids"
}
func (r *RepoConfFlowIdImpl) Create(ctx context.Context, data any) error {
	rs := r.db.Table(r.Table()).Create(data)
	return rs.Error
}
func (r *RepoConfFlowIdImpl) Delete(ctx context.Context, where any) (int64, error) {
	tx := r.db.Table(r.Table())
	rs := tx.Where(repo.ParseWhere(tx, where)).Delete(&ev.ConfFlowId{})
	return rs.RowsAffected, rs.Error
}
func (r *RepoConfFlowIdImpl) Update(ctx context.Context, where *any, data any) (int64, error) {
	tx := r.db.Table(r.Table())
	rs := tx.Where(repo.ParseWhere(tx, where)).Updates(data)
	return rs.RowsAffected, rs.Error
}
func (r *RepoConfFlowIdImpl) Select(columns string, where any, group string, having any, order string, page, pageSize int) ([]ev.ConfFlowId, error) {
	var data []ev.ConfFlowId
	tx := r.db.Table(r.Table())
	rs := repo.Select(tx, columns, where, group, having, order, page, pageSize).Find(&data)
	return data, rs.Error
}
func (r *RepoConfFlowIdImpl) SelectOne(columns string, where any, group string, having any, order string) (*ev.ConfFlowId, error) {
	var data ev.ConfFlowId
	tx := r.db.Table(r.Table())
	rs := repo.SelectOne(tx, columns, where, group, having, order).Find(&data)
	return &data, rs.Error
}
