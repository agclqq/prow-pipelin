package agg

import (
	"cmp"
	"context"
	"slices"

	"github.com/agclqq/prow-pipeline/domain/flow/entity"
)

type ConfFlowStage interface {
}

type ConfFlowStageImpl struct {
	entity *entity.ConfFlowStage
}

func NewConfFlowStageImpl(ent *entity.ConfFlowStage) ConfFlowStage {
	return &ConfFlowStageImpl{entity: ent}
}

// AddRearrange 新增并重排
func (s *ConfFlowStageImpl) AddRearrange(ctx context.Context, ent *entity.ConfFlowStage, ents []*entity.ConfFlowStage) []*entity.ConfFlowStage {
	slices.SortFunc(ents, func(e1 *entity.ConfFlowStage, e2 *entity.ConfFlowStage) int {
		return cmp.Compare(e1.OrderNum, e2.OrderNum)
	})

	//如果目标序号不存在或大于当前数组的order，则直接插在当前数组的末尾
	if ent.OrderNum == 0 || ent.OrderNum > len(ents) {
		ent.OrderNum = len(ents) + 1
		ents = append(ents, ent)
		return ents
	}

	ents = slices.Insert(ents, ent.OrderNum, ent)

	for i, v := range ents {
		v.OrderNum = i + 1
	}
	return ents
}

// MoveRearrange 移动重排
func (s *ConfFlowStageImpl) MoveRearrange(ctx context.Context, ent *entity.ConfFlowStage, ents []*entity.ConfFlowStage) []*entity.ConfFlowStage{
	ents = slices.DeleteFunc(ents, func(e *entity.ConfFlowStage) bool {
		return ent.ID == e.ID
	})
	return s.AddRearrange(ctx,ent,ents)
}
