package coreplugins

import (
	"lab2/src/plugins/routes"
	"lab2/src/middlewares"

	"github.com/gin-gonic/gin"
)

func MainRegister(router *gin.Engine)gin.IRoutes{
    UserRouter := routes.UserRouter
    userGroup := router.Group("/api/v1/users")
    UserRouter(userGroup)
    
    LabRouter := routes.LabRouter
    LabGroup := router.Group("/api/v1/lab")
    LabGroup.Use(middlewares.RateTimeMiddleware())
    LabRouter(LabGroup)

    return router
}