package repositories

import (
	"context"
	"sync"
	"your_project/domain/flow/entity"
)

type InMemoryConfFlowStageRepository struct {
	stages map[string]map[string]*entity.ConfFlowStage // flowID -> (stageID -> Stage)
	mu     sync.RWMutex
}

// NewInMemoryConfFlowStageRepository 创建新的内存仓储实例
func NewInMemoryConfFlowStageRepository() *InMemoryConfFlowStageRepository {
	return &InMemoryConfFlowStageRepository{
		stages: make(map[string]map[string]*entity.ConfFlowStage),
	}
}

// Exists 检查阶段是否存在
func (r *InMemoryConfFlowStageRepository) Exists(flowID string, stageID string) (bool, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	_, exists := r.stages[flowID][stageID]
	return exists, nil
}

// Delete 删除阶段
func (r *InMemoryConfFlowStageRepository) Delete(flowID string, stageID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.stages[flowID], stageID)
	return nil
} 