package main

import (
	"fmt"
	"time"
)

func ping(ping chan<- string, msg string) {
	ping <- msg
}

func pong(ping <-chan string, pong chan<- string) {
	msg := <-ping
	time.Sleep(time.Second)
	pong <- msg
}

func main() {
	pings := make(chan string)
	pongs := make(chan string)
	go ping(pings, "hello message")
	go pong(pings, pongs)
	fmt.Println(<-pongs)
}
