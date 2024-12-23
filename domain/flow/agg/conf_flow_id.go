package agg

import (
	"github.com/agclqq/prow-pipeline/domain/flow/entity"
)

type AggConfFlowId interface {
}
type AggConfFlowIdImpl struct {
	entity *entity.ConfFlowId
}

var _ AggConfFlowId = (*AggConfFlowIdImpl)(nil)

func NewAggConfFlowId(entity *entity.ConfFlowId) AggConfFlowId {
	return &AggConfFlowIdImpl{entity: entity}
}
