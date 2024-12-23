package controller

import (
	"github.com/agclqq/prow-framework/validator"
	"github.com/agclqq/prow-pipeline/app/http/controller/response"
	"github.com/agclqq/prow-pipeline/domain/flow/entity"
	"github.com/agclqq/prow-pipeline/domain/flow/svr"
	"github.com/agclqq/prow-pipeline/domain/flow/vo"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ConfFlowStage struct {
}

func (c ConfFlowStage) Index(ctx *gin.Context) {

}
func (c ConfFlowStage) Show(ctx *gin.Context) {

}
func (c ConfFlowStage) Store(ctx *gin.Context) {
	var vldUri vo.VldConfFlowStageStoreUri
	if err := ctx.ShouldBindUri(&vldUri); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Failure(ctx, response.INVALID_PARAMETER, validator.GetError(err).Error()))
		return
	}
	var vld vo.VldConfFlowStageStore
	if err := ctx.ShouldBindJSON(&vld); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Failure(ctx, response.INVALID_PARAMETER, validator.GetError(err).Error()))
		return
	}
	fSvr := svr.NewFlowSvrImpl()
	flow, err := fSvr.VerifyV0Flow(ctx, vldUri.FlowId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Failure(ctx, response.SERVER_ERROR, err.Error()))
		return
	}
	fsSvr := svr.NewFlowStageSvrImpl()
	err = fsSvr.VerifyStageName(ctx, flow.ID, vld.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Failure(ctx, response.SERVER_ERROR, err.Error()))
		return
	}

	data := &entity.ConfFlowStage{
		FlowId:   flow.ID,
		Name:     vld.Name,
		OrderNum: vld.Order,
	}

	err = fsSvr.CreateStage(ctx, data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Failure(ctx, response.SERVER_ERROR, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(ctx))
}
func (c ConfFlowStage) Update(ctx *gin.Context) {
	var vldUri vo.VldConfFlowStageUpdateUri
	if err := ctx.ShouldBindUri(&vldUri); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Failure(ctx, response.INVALID_PARAMETER, validator.GetError(err).Error()))
		return
	}
	var vld vo.VldConfFlowStageStore
	if err := ctx.ShouldBindJSON(&vld); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Failure(ctx, response.INVALID_PARAMETER, validator.GetError(err).Error()))
		return
	}
	fSvr := svr.NewFlowSvrImpl()
	_, err := fSvr.VerifyV0Flow(ctx, vldUri.FlowId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Failure(ctx, response.SERVER_ERROR, err.Error()))
		return
	}
	fsSvr := svr.NewFlowStageSvrImpl()
	stage, err := fsSvr.SelectOneStageById(ctx, vldUri.StageId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.Failure(ctx, response.INVALID_PARAMETER, err.Error()))
		return
	}
	if stage.FlowId != vldUri.FlowId || stage.Version != 0 {
		ctx.JSON(http.StatusNotFound, response.Failure(ctx, response.INVALID_PARAMETER, "未找到与flow配置的阶段"))
		return
	}
	data := &entity.ConfFlowStage{
		ID:       vldUri.StageId,
		FlowId:   vldUri.FlowId,
		Name:     vld.Name,
		OrderNum: vld.Order,
	}
	err = fsSvr.UpdateStage(ctx, stage, data, &entity.ConfFlowStage{ID: vldUri.StageId})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Failure(ctx, response.SERVER_ERROR, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(ctx))
}
func (c ConfFlowStage) Destroy(ctx *gin.Context) {
	var vldUri vo.VldConfFlowStageUpdateUri
	if err := ctx.ShouldBindUri(&vldUri); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Failure(ctx, response.INVALID_PARAMETER, validator.GetError(err).Error()))
		return
	}
	_, err := svr.NewFlowStageSvrImpl().DeleteStage(ctx, vldUri.FlowId, vldUri.StageId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Failure(ctx, response.SERVER_ERROR, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, response.Success(ctx))
}
