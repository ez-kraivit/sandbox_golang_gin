package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func RateTimeMiddleware()gin.HandlerFunc{
    return func(c *gin.Context) {
        fmt.Println(c.Request.Header)
        fmt.Println("MiddlewareRateTime")
        requests := make(chan int, 1)
        for i := 1; i <= 1; i++ {
          requests <- i
        }
        close(requests)
      
        limiter := time.NewTicker(1*time.Second)

        for req := range requests {
          <-limiter.C
          fmt.Println("request", req, time.Now())
        }

        c.Next()
    }
}