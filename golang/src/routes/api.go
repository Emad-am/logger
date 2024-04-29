package routes

import (
	logcontroller "logger/src/controllers/log"
	"logger/src/middlewares"

	"github.com/gin-gonic/gin"
)

func DefineMiddlewares(router *gin.Engine) {
	router.Use(middlewares.Auth())
}

func DefineRoutes(router *gin.Engine) {
	log := router.Group("/log")
	{
		log.POST("/report-error", logcontroller.ReportError)
	}

}
