package services

import (
	"your_project/internal/domain/repositories"
	"errors"
)

type ConfFlowStageService struct {
	stageRepo repositories.ConfFlowStageRepository
}

// NewConfFlowStageService 创建新的服务实例
func NewConfFlowStageService(stageRepo repositories.ConfFlowStageRepository) *ConfFlowStageService {
	return &ConfFlowStageService{stageRepo: stageRepo}
}

// DeleteStage 删除指定的阶段
func (s *ConfFlowStageService) DeleteStage(flowID string, stageID string) error {
	// 检查阶段是否存在
	exists, err := s.stageRepo.Exists(flowID, stageID)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("stage not found")
	}

	// 删除阶段
	return s.stageRepo.Delete(flowID, stageID)
} 