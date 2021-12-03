package main

import (
	"lab2/src/coreplugins"

	"github.com/gin-gonic/gin"
)

func main() {
    // Cors for cross origin allowance
    Cors := coreplugins.CorsFunc
	router := gin.Default()
    router.Use(Cors())

    // Env variables 
    EnvJSON := coreplugins.Env()

    // Register Go
    GoRegister := coreplugins.MainRegister
    GoRegister(router)
	router.Run(EnvJSON.Port)
    
}

