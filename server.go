package main

import (
    "log"
	"lab2/src/coreplugins"

	"github.com/gin-gonic/gin"
)

func main() {
    // coreplugins.CronJobFunc()
    
    // Cors for cross origin allowance
    Cors := coreplugins.CorsFunc
	router := gin.Default()
    router.Use(Cors())

    // Env variables 
    EnvJSON := coreplugins.Env()

    // Register Go
    GoRegister := coreplugins.MainRegister
    GoRegister(router)
    log.Fatal(router.Run(EnvJSON.Domain+EnvJSON.Port))
    
}

