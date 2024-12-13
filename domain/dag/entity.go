package flow

type ConfDag struct {
	ID         int `json:"id,omitempty"`
	FlowIId    int `json:"flow_iid,omitempty"`
	StageId    int `json:"stage_id,omitempty"`
	StepId     int `json:"step_id,omitempty"`
	FromStepId int `json:"from_step_id,omitempty"`
	NextStepId int `json:"next_step_id,omitempty"`
}
