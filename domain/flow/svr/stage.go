package svr

import (
	"context"
	"errors"

	"github.com/agclqq/prow-pipeline/boot"
	"github.com/agclqq/prow-pipeline/domain/flow/entity"
	"github.com/agclqq/prow-pipeline/domain/flow/repo"
	"gorm.io/gorm"
)

type FlowStageSvr interface {
	VerifyStage(ctx context.Context, data *entity.ConfFlowStage) (*entity.ConfFlowStage, error)
	UpdateStage(ctx context.Context, currentStage, newStage, where *entity.ConfFlowStage) error
	SelectOneStageById(ctx context.Context, id int) (*entity.ConfFlowStage, error)
	VerifyStageName(ctx context.Context, flowId int, stageName string) error
	SelectV0Stages(ctx context.Context, flowIId int, orderNum int) ([]*entity.ConfFlowStage, error)
	CreateStage(ctx context.Context, data *entity.ConfFlowStage) error
	DeleteStage(ctx context.Context, flowID int, stageID int) (int64, error)
}

type flowStageSvrImpl struct {
	db *gorm.DB
}

func NewFlowStageSvrImpl(db ...*gorm.DB) FlowStageSvr {
	if len(db) == 0 {
		return &flowStageSvrImpl{db: boot.GetDbW()}
	}
	return &flowStageSvrImpl{db: db[0]}
}

func (f *flowStageSvrImpl) VerifyStage(ctx context.Context, data *entity.ConfFlowStage) (*entity.ConfFlowStage, error) {
	stage, err := f.SelectOneStageById(ctx, data.ID)
	if err != nil {
		return nil, err
	}
	if stage == nil || stage.ID == 0 {
		return nil, errors.New("未找到对应的阶段")
	}
	return stage, nil
}
func (f *flowStageSvrImpl) UpdateStage(ctx context.Context, currentStage, newStage, where *entity.ConfFlowStage) error {
	tx := f.db.Begin()
	var err error
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()
	cfsRepo := repo.NewConfFlowStageImpl(f.db)
	if currentStage.OrderNum > newStage.OrderNum { //从后往前调
		stages, err := cfsRepo.Select(ctx, "", map[string]any{"flow_id": currentStage.FlowId, "version": 0, "order_num >= ?": newStage.OrderNum, "order_num < ?": currentStage.OrderNum}, "", nil, "", 0, 0)
		if err != nil {
			return err
		}
		ids := make([]int, len(stages))
		for i, v := range stages {
			ids[i] = v.ID
		}
		rowNum, err := cfsRepo.Update(ctx, map[string]any{"id in ?": ids}, map[string]any{"order_num": gorm.Expr("order_num + ?", 1)})
		if err != nil {
			return err
		}
		if int(rowNum) > len(stages) {
			return errors.New("更新失败")
		}
	}
	if currentStage.OrderNum < newStage.OrderNum { //从前往后调
		stages, err := cfsRepo.Select(ctx, "", map[string]any{"flow_id": currentStage.FlowId, "version": 0, "order_num > ?": currentStage.OrderNum, "order_num < ?": newStage.OrderNum}, "", nil, "", 0, 0)
		if err != nil {
			return err
		}
		ids := make([]int, len(stages))
		for i, v := range stages {
			ids[i] = v.ID
		}
		rowNum, err := cfsRepo.Update(ctx, map[string]any{"id in ?": ids}, map[string]any{"order_num": gorm.Expr("order_num + ?", 1)})
		if err != nil {
			return err
		}
		if int(rowNum) > len(stages) {
			return errors.New("更新失败")
		}
	}
	_, err = cfsRepo.Update(ctx, map[string]any{"id": currentStage.ID}, newStage)
	if err != nil {
		return err
	}
	return nil
}
func (f *flowStageSvrImpl) SelectOneStageById(ctx context.Context, id int) (*entity.ConfFlowStage, error) {
	return repo.NewConfFlowStageImpl(f.db).SelectOne(ctx, "", &entity.ConfFlowStage{ID: id}, "", nil, "")
}
func (f *flowStageSvrImpl) VerifyStageName(ctx context.Context, flowId int, stageName string) error {
	stages, err := repo.NewConfFlowStageImpl(f.db).Select(ctx, "id", map[string]any{"flow_id": flowId, "version": 0, "name": stageName}, "", nil, "", 0, 0)
	if err != nil {
		return err
	}
	if len(stages) > 0 {
		return errors.New("阶段名称重复")
	}
	return nil
}
func (f *flowStageSvrImpl) SelectV0Stages(ctx context.Context, flowId int) ([]*entity.ConfFlowStage, error) {
	return repo.NewConfFlowStageImpl(f.db).Select(ctx, "", map[string]any{"flow_id": flowId, "version": 0}, "", nil, "order_num asc", 0, 0)
}

func (f *flowStageSvrImpl) CreateStage(ctx context.Context, data *entity.ConfFlowStage) error {
	var err error
	stages, err := f.SelectV0Stages(ctx, data.FlowId)
	if err != nil {
		return err
	}

	ids := make([]int, len(stages))
	for i := range stages {
		ids[i] = stages[i].ID
	}
	tx := f.db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()
	cfs := repo.NewConfFlowStageImpl(tx)
	rowNum, err := cfs.Update(ctx, map[string]any{"id in ?": ids}, map[string]any{"order_num": gorm.Expr("order_num + ?", 1)})
	if err != nil {
		return err
	}
	if int(rowNum) != len(stages) {
		return errors.New("批量变更stage序号失败")
	}
	return cfs.Create(ctx, data)
}
func (f *flowStageSvrImpl) DeleteStage(ctx context.Context, flowID int, stageID int) (int64, error) {
	// 删除阶段
	where := &entity.ConfFlowStage{FlowId: flowID, ID: stageID}
	return repo.NewConfFlowStageImpl(f.db).Delete(ctx, where)
}
