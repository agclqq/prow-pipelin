package fty

import (
	"github.com/agclqq/prow-pipeline/boot"
	"github.com/agclqq/prow-pipeline/domain/flow/svr"
)

func NewFlowSvrImplFty() svr.FlowSvr {
	return svr.NewFlowSvrImpl(boot.GetDbW())
}
