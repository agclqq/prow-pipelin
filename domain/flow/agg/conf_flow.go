package agg

import (
	"github.com/agclqq/prow-pipeline/domain/flow/entity"
)

type ConfFlow interface {
	Empty() bool
}
type confFlowImpl struct {
	entity *entity.ConfFlow
}

var _ ConfFlow = (*confFlowImpl)(nil)

func NewAggConfFlowImpl(entity *entity.ConfFlow) ConfFlow {
	return &confFlowImpl{entity: entity}
}

func (a *confFlowImpl) Empty() bool {
	if a.entity == nil || a.entity.ID == 0 {
		return true
	}
	return false
}
