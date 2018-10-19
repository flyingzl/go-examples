// [定时器](../timers/) 是当你想要在未来某一刻执行一次时
// 使用的 - _打点器_ 则是当你想要在固定的时间间隔重复执行
// 准备的。这里是一个打点器的例子，它将定时的执行，直到我
// 们将它停止。

package main

import "time"
import "fmt"

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main() {

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for i := 1; i <= 9; i++ {
		jobs <- i
	}

	close(jobs)

	for a := 1; a <= 9; a++ {
		<-results
	}
}
