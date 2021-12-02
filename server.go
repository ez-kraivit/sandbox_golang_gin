package main

import (
    "lab2/src/coreplugins"

    "github.com/gin-gonic/gin"
)

func main() {

    var Cors = coreplugins.CorsFunc
	router := gin.Default()
    router.Use(Cors())
    
    var GoRegister = coreplugins.MainRegister
    GoRegister(router)
    
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
