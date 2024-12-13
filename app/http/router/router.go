package router

import (
	"github.com/agclqq/prow-framework/http/restful/router"
	"github.com/gin-gonic/gin"

	"github.com/agclqq/prow-pipeline/app/http/controller"
)

func Register(eng *gin.Engine) {
	apiGroup := eng.Group("/api")
	{
		router.ApiResource(apiGroup, "/demo", &controller.Demo{})
		router.ApiResource(apiGroup, "/flowId", &controller.ConfFlowId{}, router.STORE)
		router.ApiResource(apiGroup, "/flow", &controller.ConfFlow{})
	}
}
