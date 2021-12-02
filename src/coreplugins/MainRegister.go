package coreplugins

import (
	"lab2/src/plugins/routes"

	"github.com/gin-gonic/gin"
)

func MainRegister(router *gin.Engine)gin.IRoutes{
    var UserRouter = routes.UserRouter
    userGroup := router.Group("/api/v1/users")
    UserRouter(userGroup)
    return router
}