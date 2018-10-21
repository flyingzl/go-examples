package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second)
		c1 <- "hello"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "world"
	}()

	select {
	case msg := <-c1:
		fmt.Println("c1: ", msg)
	}

	select {
	case msg := <-c2:
		fmt.Println("c2: ", msg)
	case <-time.After(time.Second):
		fmt.Println("c2 timeout")
	}
}
