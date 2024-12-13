package controller

import (
	"github.com/agclqq/prow-framework/validator"
	"github.com/agclqq/prow-pipeline/app/http/controller/response"
	"github.com/agclqq/prow-pipeline/boot"
	"github.com/agclqq/prow-pipeline/domain/flow/agg/ev"
	"github.com/agclqq/prow-pipeline/domain/flow/svr"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ConfFlow struct {
}

func (f ConfFlow) Index(ctx *gin.Context) {
	var vld ev.VldFlowIndex
	err := ctx.ShouldBind(&vld)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Failure(ctx, response.INVALID_PARAMETER, validator.GetError(err).Error()))
		return
	}
	where := &ev.ConfFlow{
		ID:       vld.Id,
		FlowId:   vld.FlowId,
		Version:  vld.Version,
		Modifier: vld.Modifier,
	}
	total, list, err := svr.NewFlowSvrImpl().PaginationFlow(ctx, "", where, "", nil, "id desc", vld.Page, vld.PageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Failure(ctx, response.SERVER_ERROR, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(ctx, map[string]any{"total": total, "list": list}))
}
func (f ConfFlow) Show(ctx *gin.Context) {

}
func (f ConfFlow) Store(ctx *gin.Context) {
	var vld ev.VldFlowPost
	err := ctx.ShouldBind(&vld)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Failure(ctx, response.INVALID_PARAMETER, validator.GetError(err).Error()))
		return
	}
	tx := boot.GetDbW().Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()
	fs := svr.NewFlowSvrImpl(tx)
	flowIdData := &ev.ConfFlowId{}
	err = fs.CreateFlowId(ctx, flowIdData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Failure(ctx, response.SERVER_ERROR, err.Error()))
		return
	}
	flowData := &ev.ConfFlow{
		FlowId:           flowIdData.ID,
		Name:             vld.Name,
		ParallelNum:      vld.ParallelNum,
		ParallelStrategy: vld.ParallelStrategy,
		ResourceId:       vld.ResourceId,
		ResourceConf:     vld.ResourceConf,
		BeforeRun:        vld.BeforeRun,
		AfterRun:         vld.AfterRun,
		Modifier:         vld.Modifier,
	}
	err = fs.CreateFlow(ctx, flowData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Failure(ctx, response.SERVER_ERROR, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(ctx, flowData.ID))
}
func (f ConfFlow) Update(ctx *gin.Context) {
	var vld ev.VldFlowUpdate
	err := ctx.ShouldBind(&vld)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Failure(ctx, response.INVALID_PARAMETER, validator.GetError(err).Error()))
		return
	}
	data := &ev.ConfFlow{
		Name:             vld.Name,
		ParallelNum:      vld.ParallelNum,
		ParallelStrategy: vld.ParallelStrategy,
		ResourceId:       vld.ResourceId,
		ResourceConf:     vld.ResourceConf,
		BeforeRun:        vld.BeforeRun,
		AfterRun:         vld.AfterRun,
		Modifier:         vld.Modifier,
	}
	_, err = svr.NewFlowSvrImpl().UpdateFlow(ctx, &ev.ConfFlow{ID: vld.Id, FlowId: vld.FlowId}, data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Failure(ctx, response.SERVER_ERROR, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(ctx))
}
func (f ConfFlow) Destroy(ctx *gin.Context) {

}
