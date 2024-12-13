package controller

import (
	"github.com/agclqq/prow-pipeline/app/http/controller/response"
	"github.com/agclqq/prow-pipeline/domain/flow/agg/ev"
	"github.com/agclqq/prow-pipeline/domain/flow/svr"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ConfFlowId struct {
}

func (c ConfFlowId) Index(ctx *gin.Context) {

}
func (c ConfFlowId) Show(ctx *gin.Context) {

}
func (c ConfFlowId) Store(ctx *gin.Context) {
	data := &ev.ConfFlowId{}
	err := svr.NewFlowSvrImpl().CreateFlowId(ctx, data)
	if err != nil {
		ctx.JSON(http.StatusOK, response.Failure(ctx, response.SERVER_ERROR, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(ctx, data.ID))
}
func (c ConfFlowId) Update(ctx *gin.Context) {

}
func (c ConfFlowId) Destroy(ctx *gin.Context) {

}
