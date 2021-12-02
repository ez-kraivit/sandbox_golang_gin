package coreplugins

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsFunc()gin.HandlerFunc{
    return cors.New(cors.Config{
        AllowOrigins:     []string{"https://kraivit.com"},
        AllowMethods:     []string{"PUT", "PATCH","GET","POST","DELETE"},
        AllowHeaders:     []string{"Origin"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        AllowOriginFunc: func(origin string) bool {
          return origin == "https://kraivit.com"
        },
    })
}