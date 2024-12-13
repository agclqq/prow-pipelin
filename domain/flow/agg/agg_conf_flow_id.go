package agg

import (
	"github.com/agclqq/prow-pipeline/domain/flow/agg/ev"
)

type AggConfFlowId interface {
}
type AggConfFlowIdImpl struct {
	entity *ev.ConfFlowId
}

var _ AggConfFlowId = (*AggConfFlowIdImpl)(nil)

func NewAggConfFlowId(entity *ev.ConfFlowId) AggConfFlowId {
	return &AggConfFlowIdImpl{entity: entity}
}
