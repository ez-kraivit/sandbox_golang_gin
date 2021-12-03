package routes

import (
    "lab2/src/plugins/controllers"
    "github.com/gin-gonic/gin"
)

func UserRouter(router *gin.RouterGroup)gin.IRoutes{
    return router.GET("", controllers.UserController)
}