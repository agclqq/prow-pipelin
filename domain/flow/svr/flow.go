package svr

import (
	"context"
	"errors"
	"github.com/agclqq/prow-pipeline/boot"
	"github.com/agclqq/prow-pipeline/domain/flow/agg"
	"github.com/agclqq/prow-pipeline/domain/flow/entity"
	"github.com/agclqq/prow-pipeline/domain/flow/repo"
	"gorm.io/gorm"
)

type FlowSvr interface {
	CreateFlowId(ctx context.Context, data any) error
	CreateFlow(ctx context.Context, data any) error
	UpdateFlow(ctx context.Context, where any, data any) (int64, error)
	PaginationFlow(ctx context.Context, columns string, where any, group string, having any, order string, page, pageSize int) (int64, []*entity.ConfFlow, error)
	SelectOneV0Flow(ctx context.Context, flowId int) (*entity.ConfFlow, error)
	VerifyV0Flow(ctx context.Context, flowId int) (*entity.ConfFlow, error)
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
	return repo.NewConfFlowIdImpl(f.db).Create(ctx, data)
}

func (f *flowSvrImpl) CreateFlow(ctx context.Context, data any) error {
	return repo.NewConfFlowImpl(f.db).Create(ctx, data)
}

func (f *flowSvrImpl) SelectOneV0Flow(ctx context.Context, flowId int) (*entity.ConfFlow, error) {
	return repo.NewConfFlowImpl(f.db).SelectOne(ctx, "id", map[string]any{"flow_id": flowId, "version": 0}, "", nil, "")
}

func (f *flowSvrImpl) UpdateFlow(ctx context.Context, where any, data any) (int64, error) {
	return repo.NewConfFlowImpl(f.db).Update(ctx, where, data)
}
func (f *flowSvrImpl) PaginationFlow(ctx context.Context, columns string, where any, group string, having any, order string, page, pageSize int) (int64, []*entity.ConfFlow, error) {
	if order == "" {
		order = "id desc"
	}
	return repo.NewConfFlowImpl(f.db).Pagination(ctx, columns, where, group, having, order, page, pageSize)
}

func (f *flowSvrImpl) VerifyV0Flow(ctx context.Context, flowId int) (*entity.ConfFlow, error) {
	flow, err := f.SelectOneV0Flow(ctx, flowId)
	if err != nil {
		return nil, err
	}
	if agg.NewAggConfFlowImpl(flow).Empty() {
		return nil, errors.New("未找到流水线")
	}
	return flow, nil
}
