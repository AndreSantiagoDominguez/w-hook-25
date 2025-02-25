package infrastructure

import (
	infrastructure "web-hook/pull_reques_webhook/infrastructure/controller"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(e *gin.Engine) {
	routesPull := e.Group("/pullRequest") 
	{
		routesPull.POST("/proccess", infrastructure.HandlePullRequestEvent)
	}
}