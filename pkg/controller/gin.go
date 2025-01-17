package controller

import (
	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	handlerServer HandlerServer
}

func NewHttpService(handlerServer HandlerServer) *HttpServer {
	return &HttpServer{
		handlerServer: handlerServer,
	}
}

func (h *HttpServer) StartHttpServer() {
	r := gin.Default()
	//create pipeline job
	r.POST("/pipeline", h.handlerServer.createPipeline)
	//update pipeline job
	r.PUT("/pipeline/:oldName", h.handlerServer.updatePipeline)
	//get pipeline job
	r.GET("/pipeline/:name", h.handlerServer.getPipeline)
	//delete pipeline job and pipeline job detail
	r.DELETE("/pipeline/:name", h.handlerServer.deletePipeline)
	//get pipeline job list
	r.GET("/pipeline", h.handlerServer.pipelineList)
	//get pipeline job detail info
	r.GET("/pipeline/:name/detail/:id", h.handlerServer.getPipelineDetail)
	//delete pipeline job detail
	r.DELETE("/pipeline/:name/detail/:id", h.handlerServer.deleteJobDetail)
	//get pipeline job detail list
	r.GET("/pipeline/:name/details", h.handlerServer.getPipelineDetailList)
	//exec pipeline job
	r.POST("/pipeline/:name/exec", h.handlerServer.execPipeline)
	//re exec pipeline detail job
	r.POST("/pipeline/:name/:id/re-exec", h.handlerServer.reExecuteJob)
	//stop pipeline job
	r.POST("/pipeline/:name/:id/stop", h.handlerServer.stopJobDetail)
	r.GET("/pipeline/:name/logs/:id", h.handlerServer.getJobLog)
	r.GET("/pipeline/:name/logs/:id/:stagename", h.handlerServer.getJobStageLog)
	// get template list
	r.GET("/pipeline/templates", h.handlerServer.getTemplates)
	r.GET("/pipeline/template/:id", h.handlerServer.getTemplateDetail)
	r.GET("/ping", func(c *gin.Context) {
		//输出json结果给调用方
		Success("", c)
	})
	r.Run(":8080") // listen and serve on
}
