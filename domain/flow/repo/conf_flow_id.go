package repo

import (
	"context"
	"github.com/agclqq/prow-framework/db/repo"
	"github.com/agclqq/prow-pipeline/domain/flow/entity"
	"gorm.io/gorm"
)

type ConfFlowId interface {
	Create(ctx context.Context, data any) error
	Delete(ctx context.Context, where any) (int64, error)
	Update(ctx context.Context, where *any, data any) (int64, error)
	SelectOne(columns string, where any, group string, having any, order string) (*entity.ConfFlowId, error)
}
type ConfFlowIdImpl struct {
	db *gorm.DB
}

func NewConfFlowIdImpl(db *gorm.DB) ConfFlowId {
	return &ConfFlowIdImpl{db: db}
}

func (r *ConfFlowIdImpl) Table() string {
	return "conf_flow_ids"
}
func (r *ConfFlowIdImpl) Create(ctx context.Context, data any) error {
	rs := r.db.Table(r.Table()).Create(data)
	return rs.Error
}
func (r *ConfFlowIdImpl) Delete(ctx context.Context, where any) (int64, error) {
	tx := r.db.Table(r.Table())
	rs := tx.Where(repo.ParseWhere(tx, where)).Delete(&entity.ConfFlowId{})
	return rs.RowsAffected, rs.Error
}
func (r *ConfFlowIdImpl) Update(ctx context.Context, where *any, data any) (int64, error) {
	tx := r.db.Table(r.Table())
	rs := tx.Where(repo.ParseWhere(tx, where)).Updates(data)
	return rs.RowsAffected, rs.Error
}
func (r *ConfFlowIdImpl) Select(columns string, where any, group string, having any, order string, page, pageSize int) ([]entity.ConfFlowId, error) {
	var data []entity.ConfFlowId
	tx := r.db.Table(r.Table())
	rs := repo.Select(tx, columns, where, group, having, order, page, pageSize).Find(&data)
	return data, rs.Error
}
func (r *ConfFlowIdImpl) SelectOne(columns string, where any, group string, having any, order string) (*entity.ConfFlowId, error) {
	var data entity.ConfFlowId
	tx := r.db.Table(r.Table())
	rs := repo.SelectOne(tx, columns, where, group, having, order).Find(&data)
	return &data, rs.Error
}
