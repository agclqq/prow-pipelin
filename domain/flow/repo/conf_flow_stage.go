package repo

import (
	"context"
	"github.com/agclqq/prow-framework/db/repo"
	"github.com/agclqq/prow-pipeline/domain/flow/entity"
	"gorm.io/gorm"
)

type ConfFlowStage interface {
	Create(ctx context.Context, data any) error
	Delete(ctx context.Context, where any) (int64, error)
	Update(ctx context.Context, where any, data any) (int64, error)
	Select(ctx context.Context, columns string, where any, group string, having any, order string, page, pageSize int) ([]*entity.ConfFlowStage, error)
	SelectOne(ctx context.Context, columns string, where any, group string, having any, order string) (*entity.ConfFlowStage, error)
}
type ConfFlowStageImpl struct {
	db *gorm.DB
}

func NewConfFlowStageImpl(db *gorm.DB) ConfFlowStage {
	return &ConfFlowStageImpl{db: db}
}

func (r *ConfFlowStageImpl) Table() string {
	return "conf_flow_stages"
}
func (r *ConfFlowStageImpl) Create(ctx context.Context, data any) error {
	rs := r.db.WithContext(ctx).Table(r.Table()).Create(data)
	return rs.Error
}
func (r *ConfFlowStageImpl) Delete(ctx context.Context, where any) (int64, error) {
	tx := r.db.WithContext(ctx).Table(r.Table())
	rs := tx.Where(repo.ParseWhere(tx, where)).Delete(&entity.ConfFlow{})
	return rs.RowsAffected, rs.Error
}
func (r *ConfFlowStageImpl) Update(ctx context.Context, where any, data any) (int64, error) {
	tx := r.db.WithContext(ctx).Table(r.Table())
	tx = repo.ParseWhere(tx, where)
	rs := tx.Updates(data)
	return rs.RowsAffected, rs.Error
}

func (r *ConfFlowStageImpl) Select(ctx context.Context, columns string, where any, group string, having any, order string, page, pageSize int) ([]*entity.ConfFlowStage, error) {
	var data []*entity.ConfFlowStage
	tx := r.db.WithContext(ctx).Table(r.Table())
	rs := repo.Select(tx, columns, where, group, having, order, page, pageSize).Find(&data)
	return data, rs.Error
}
func (r *ConfFlowStageImpl) SelectOne(ctx context.Context, columns string, where any, group string, having any, order string) (*entity.ConfFlowStage, error) {
	var data entity.ConfFlowStage
	tx := r.db.WithContext(ctx).Table(r.Table())
	rs := repo.SelectOne(tx, columns, where, group, having, order).Find(&data)
	return &data, rs.Error
}
