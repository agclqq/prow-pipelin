package svr

import (
	"context"
	"github.com/agclqq/prow-pipeline/boot"
	"github.com/agclqq/prow-pipeline/domain/flow/agg/ev"
	"github.com/agclqq/prow-pipeline/domain/flow/repo"
	"gorm.io/gorm"
)

type FlowSvr interface {
	CreateFlowId(ctx context.Context, data any) error
	CreateFlow(ctx context.Context, data any) error
	UpdateFlow(ctx context.Context, where any, data any) (int64, error)
	PaginationFlow(ctx context.Context, columns string, where any, group string, having any, order string, page, pageSize int) (int64, []ev.ConfFlow, error)
}

type flowSvrImpl struct {
	db *gorm.DB
}

func NewFlowSvrImpl(db ...*gorm.DB) FlowSvr {
	if len(db) == 0 {
		return &flowSvrImpl{db: boot.GetDbW()}
	}
	return &flowSvrImpl{db: db[0]}
}

func (f *flowSvrImpl) CreateFlowId(ctx context.Context, data any) error {
	return repo.NewRepoConfFlowIdImpl(f.db).Create(ctx, data)
}

func (f *flowSvrImpl) CreateFlow(ctx context.Context, data any) error {
	return repo.NewConfFlow(f.db).Create(ctx, data)
}

func (f *flowSvrImpl) UpdateFlow(ctx context.Context, where any, data any) (int64, error) {
	return repo.NewConfFlow(f.db).Update(ctx, where, data)
}
func (f *flowSvrImpl) PaginationFlow(ctx context.Context, columns string, where any, group string, having any, order string, page, pageSize int) (int64, []ev.ConfFlow, error) {
	if order == "" {
		order = "id desc"
	}
	return repo.NewConfFlow(f.db).Pagination(ctx, columns, where, group, having, order, page, pageSize)
}
