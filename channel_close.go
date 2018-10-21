package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	// duration为纳秒
	duration := time.Second.Nanoseconds()

	go func() {
		for {
			if val, more := <-jobs; more {
				fmt.Println("receive job ", val)
				time.Sleep(time.Duration(rand.Int63n(duration)))
			} else {
				fmt.Println("receive all jobs ")
				done <- true
				return
			}
		}
	}()

	for i := 1; i < 4; i++ {
		jobs <- i
		fmt.Println("send job ", i)
	}

	close(jobs)

	<-done
}
