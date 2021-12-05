package controllers

import (
	"fmt"
  "sync"
  "time"

	"lab2/src/types"
	"github.com/chebyrash/promise"
	"github.com/gin-gonic/gin"
)

func LabControllerDemo(c *gin.Context) {

  c.JSON(200,"Lab")
}

func LabControllerArray(c *gin.Context) {

  var name = [3]string{"Chaiyarin", "Atikom", "Kritsana"};
  fmt.Println(len(name));
  c.JSON(200,name)
}

func LabControllerMap(c *gin.Context) {

  name := &types.LabMepData{
    Name: "Chaiyarin",
    Age: 25,
    Status: true,
  }
  c.JSON(200,name)
}

func LabControllerPromise(c *gin.Context) {

  var p1 = promise.Resolve(123)
  results, _ := p1.Await()
  c.JSON(200,results)
}

func GoDate(year, month, day int) time.Time {
  return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func LabControllerTime(c *gin.Context) {
  t1 := GoDate(2022, 12, 5)
  start := time.Now()
  t := time.Now()
  elapsed := t.Sub(start)
  fmt.Println(elapsed)
  fmt.Println(t1)
  c.JSON(200,t1)
}

func LabControllerRateLimite(c *gin.Context) {
  c.JSON(200,"Rate Limite")
}

func LabControllerRace(c *gin.Context) {
  p1 := promise.New(func(resolve func(promise.Any),reject func(error)){
    time.Sleep(time.Second * 5)
    resolve("5s")
  })
  p2 :=  promise.New(func(resolve func(promise.Any),reject func(error)){
    time.Sleep(time.Second * 2)
    resolve("2s")
  })
  p3 , _ := promise.Race(p1,p2).Await()
  c.JSON(200,p3)
}

func LabControllerWorkerPool(c *gin.Context) {
  const numJobs = 10
  jobs := make(chan int, numJobs)
  results := make(chan int,numJobs)
  for w:= 1; w<= 3; w++{
    go func(w int, jobs chan int, results chan int) {
      for j := range jobs {
        fmt.Println("worker",w,"started job",j)
        time.Sleep(time.Second)
        fmt.Println("worker",w,"finished job",j)
        results <- j * 2 
      }
    }(w,jobs,results)
  }
  for j := 1; j<= numJobs; j++ {
    jobs <- j
  }
  close(jobs)
  for a := 1; a<= numJobs; a++{
    <- results
  }
  c.JSON(200,"Success")
}

func LabControllerRecovery(c *gin.Context) {
  defer func(){
    if err := recover(); err != nil {
      fmt.Printf("Recovery: %v\n", err)
      c.JSON(200,"Recovery")
    }
  }()
  panic("a problem")
}

func LabControllerPanic(c *gin.Context) {
  panic("a problem")
}

func LabControllerContext(c *gin.Context) {
  var wg sync.WaitGroup
  data := []string{"a","b","c","d","e"}
  for _ , u := range data{
    time.Sleep(time.Second)
    wg.Add(1)
    go func(u string){
      fmt.Println("Done",time.Now())
      defer wg.Done()
    }(u)
  }
  wg.Wait()
  c.JSON(200,"finish")
}