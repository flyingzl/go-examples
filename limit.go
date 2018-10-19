// [定时器](../timers/) 是当你想要在未来某一刻执行一次时
// 使用的 - _打点器_ 则是当你想要在固定的时间间隔重复执行
// 准备的。这里是一个打点器的例子，它将定时的执行，直到我
// 们将它停止。

package main

import "fmt"

import "time"

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}

	close(requests)

	limitTicket := time.Tick(time.Millisecond * 1000)

	for req := range requests {
		<-limitTicket
		fmt.Println(req)
	}

	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(time.Millisecond * 1000) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}

}
