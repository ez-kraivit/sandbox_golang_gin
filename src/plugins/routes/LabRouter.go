package routes

import (
    "lab2/src/plugins/controllers"
    "github.com/gin-gonic/gin"
)

func LabRouter(router *gin.RouterGroup)gin.IRoutes{
    router.GET("/demo", controllers.LabControllerDemo)
    router.GET("/array", controllers.LabControllerArray)
    router.GET("/map", controllers.LabControllerMap)
    router.GET("/promise", controllers.LabControllerPromise)
    router.GET("/time", controllers.LabControllerTime)
    router.GET("/ratelimite", controllers.LabControllerRateLimite)
    router.GET("/promiserace",controllers.LabControllerRace)
    router.GET("/worker",controllers.LabControllerWorkerPool)
    router.GET("/recover",controllers.LabControllerRecovery)
    router.GET("/panic",controllers.LabControllerPanic)
    router.GET("/context",controllers.LabControllerContext)
    return router
}

