package main

import (
	"fmt"
	"time"
)

func main() {
   start := time.Now()
   for i := 1; i <= 10000000; i++ {
      fmt.Println(i)
   }
   end := time.Now()
   time := time.Now().Sub(start)

   fmt.Println("시작: ", start.Format("2006-01-02 15:04:05"))
   fmt.Println("종료: ", end.Format("2006-01-02 15:04:05"))
   fmt.Println("소요시간: ", time)

}