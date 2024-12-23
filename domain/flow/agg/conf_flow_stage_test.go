package agg

import (
	"context"
	"github.com/agclqq/prow-pipeline/domain/flow/entity"
	"testing"
)

func TestConfFlowStageImpl_AddRearrange(t *testing.T) {
	type fields struct {
		entity *entity.ConfFlowStage
	}
	type args struct {
		ctx  context.Context
		ent  *entity.ConfFlowStage
		ents []*entity.ConfFlowStage
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []*entity.ConfFlowStage
	}{
		{
			name:   "t1",
			fields: fields{entity: &entity.ConfFlowStage{}},
			args: args{
				ctx: nil,
				ent: &entity.ConfFlowStage{FlowId: 5, OrderNum: 2},
				ents: []*entity.ConfFlowStage{
					{FlowId: 1, OrderNum: 1},
					{FlowId: 2, OrderNum: 4},
					{FlowId: 3, OrderNum: 3},
					{FlowId: 4, OrderNum: 2},
				},
			}, want: []*entity.ConfFlowStage{
				{FlowId: 1, OrderNum: 1},
				{FlowId: 5, OrderNum: 2},
				{FlowId: 4, OrderNum: 3},
				{FlowId: 3, OrderNum: 4},
				{FlowId: 2, OrderNum: 5},
			},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ConfFlowStageImpl{
				entity: tt.fields.entity,
			}
			got := s.AddRearrange(tt.args.ctx, tt.args.ent, tt.args.ents)
			if len(got) != len(tt.want) {
				t.Errorf("AddRearrange() got = %v, want %v", got, tt.want)
				return
			}
			for i, v := range got {
				if v.OrderNum != tt.want[i].OrderNum {
					t.Errorf("得到的结果与期望的不一致got = %v, want %v", got, tt.want)
				}
				return
			}
		})
	}
}

func TestConfFlowStageImpl_MoveRearrange(t *testing.T) {
	type fields struct {
		entity *entity.ConfFlowStage
	}
	type args struct {
		ctx  context.Context
		ent  *entity.ConfFlowStage
		ents []*entity.ConfFlowStage
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []*entity.ConfFlowStage
	}{
		{
			name:   "test1",
			fields: fields{entity: &entity.ConfFlowStage{}},
			args: args{
				ctx: nil,
				ent: &entity.ConfFlowStage{ID: 3, OrderNum: 2},
				ents: []*entity.ConfFlowStage{
					{ID: 1, OrderNum: 1},
					{ID: 5, OrderNum: 2},
					{ID: 4, OrderNum: 3},
					{ID: 3, OrderNum: 4},
					{ID: 2, OrderNum: 5},
				},
			},
			want: []*entity.ConfFlowStage{
				{ID: 1, OrderNum: 1},
				{ID: 3, OrderNum: 2},
				{ID: 5, OrderNum: 3},
				{ID: 4, OrderNum: 4},
				{ID: 2, OrderNum: 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ConfFlowStageImpl{
				entity: tt.fields.entity,
			}
			got := s.MoveRearrange(tt.args.ctx, tt.args.ent, tt.args.ents)
			if len(got) != len(tt.want) {
				t.Errorf("AddRearrange() got = %+v, want %+v", got, tt.want)
				return
			}
			for i, v := range got {
				if v.OrderNum != tt.want[i].OrderNum {
					t.Errorf("得到的结果与期望的不一致got = %v, want %v", got, tt.want)
				}
				return
			}
		})