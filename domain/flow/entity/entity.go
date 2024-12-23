package entity

import "time"

type ConfFlowId struct {
	ID        int       `json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
type ConfFlow struct {
	ID               int       `json:"flow_iid,omitempty"`
	FlowId           int       `json:"flow_id,omitempty" binding:"required"`
	Version          int       `json:"version,omitempty"`
	Name             string    `json:"name,omitempty"`
	ParallelNum      int       `json:"parallel_num,omitempty"`
	ParallelStrategy int       `json:"parallel_strategy,omitempty"`
	ResourceId       int       `json:"resource_id,omitempty"`
	ResourceConf     string    `json:"resource_conf,omitempty"`
	BeforeRun        string    `json:"before_run,omitempty"`
	AfterRun         string    `json:"after_run,omitempty"`
	Modifier         int       `json:"modifier,omitempty"`
	CreatedAt        time.Time `json:"created_at,omitempty"`
	UpdatedAt        time.Time `json:"updated_at,omitempty"`
}

type ConfFlowDag struct {
	ID         int       `json:"id,omitempty"`
	FlowId     int       `json:"flow_id,omitempty"`
	Version    int       `json:"version,omitempty"`
	StageId    int       `json:"stage_id,omitempty"`
	StepId     int       `json:"step_id,omitempty"`
	FromStepId int       `json:"from_step_id,omitempty"`
	NextStepId int       `json:"next_step_id,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}
type ConfFlowStage struct {
	ID        int       `json:"id,omitempty"`
	FlowId    int       `json:"flow_id,omitempty" `
	Version   int       `json:"version,omitempty" `
	Name      string    `json:"name,omitempty" binding:"required" alias:"阶段名称"`
	OrderNum  int       `json:"order_num,omitempty"  binding:"required" alias:"顺序"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
type ConfFlowAtom struct {
	ID           int       `json:"id,omitempty"`
	FlowId       int       `json:"flow_id,omitempty"`
	Version      int       `json:"version,omitempty"`
	Name         string    `json:"name,omitempty"`
	Type         int       `json:"type,omitempty"`
	ResourceId   int       `json:"resource_id,omitempty"`
	ResourceConf string    `json:"resource_conf,omitempty"`
	BeforeRun    string    `json:"before_run,omitempty"`
	AfterRun     string    `json:"after_run,omitempty"`
	Image        string    `json:"image,omitempty"`
	Run          string    `json:"run,omitempty"`
	Modifier     int       `json:"modifier,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
}
