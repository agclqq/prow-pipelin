package ev

type VldFlowPost struct {
	FlowId           int    `json:"flow_id"`
	Name             string `json:"name" binding:"required"`
	ParallelNum      int    `json:"parallel_num,default=1"`
	ParallelStrategy int    `json:"parallel_strategy,default=1"`
	ResourceId       int    `json:"resource_id"`
	ResourceConf     string `json:"resource_conf"`
	BeforeRun        string `json:"before_run"`
	AfterRun         string `json:"after_run"`
	Modifier         int    `json:"modifier"`
}

type VldFlowUpdate struct {
	Id               int    `uri:"flow" binding:"required"`
	FlowId           int    `json:"flow_id" binding:"required"`
	Name             string `json:"name"`
	ParallelNum      int    `json:"parallel_num,default=1"`
	ParallelStrategy int    `json:"parallel_strategy,default=1"`
	ResourceId       int    `json:"resource_id"`
	ResourceConf     string `json:"resource_conf"`
	BeforeRun        string `json:"before_run"`
	AfterRun         string `json:"after_run"`
	Modifier         int    `json:"modifier"`
}
type VldFlowIndex struct {
	Id       int `json:"flow"`
	FlowId   int `json:"flow_id"`
	Version  int `json:"version"`
	Modifier int `json:"modifier"`
	Page     int `json:"page,default=1"`
	PageSize int `json:"page_size,default=10"`
}
