package main

import (
	"fmt"
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
    fmt.Println("Go server is running on ",EnvJSON.Domain+EnvJSON.Port)
    router.Run(EnvJSON.Port)
    
}

