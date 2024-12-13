package agg

import (
	"github.com/agclqq/prow-pipeline/domain/flow/agg/ev"
)

type AggConfFlow interface {
}
type AggConfFlowImpl struct {
	entity *ev.ConfFlow
}

var _ AggConfFlow = (*AggConfFlowImpl)(nil)

func NewAggConfFlowImpl(entity *ev.ConfFlow) AggConfFlow {
	return &AggConfFlowImpl{entity: entity}
}
