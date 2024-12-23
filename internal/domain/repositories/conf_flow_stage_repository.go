package repositories

type ConfFlowStageRepository interface {
    Exists(flowID string, stageID string) (bool, error)
    Delete(flowID string, stageID string) error
} 