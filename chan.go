package main

import "fmt"
import "time"

func main() {
	message := make(chan string)
	go func() {
		fmt.Println("love")
		time.Sleep(time.Second)
		message <- "hello, baby"
	}()

	msg := <-message
	fmt.Println("done, and the message is: ", msg)

}
