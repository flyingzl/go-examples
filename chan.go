package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func getRandom() int {
	c := make(chan int)
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		go func() {
			c <- int(rand.Int31n(1000))
		}()
	}

	return <-c
}

func main() {
	message := make(chan string)
	go func() {
		fmt.Println("love")
		time.Sleep(time.Second)
		message <- "hello, baby"
	}()

	msg := <-message
	fmt.Println("done, and the message is: ", msg)

	z := getRandom()
	fmt.Println(z)

	pool := &sync.Pool{
		New: func() interface{} {
			return 10
		},
	}

	fmt.Println(pool.Get().(int))
	pool.Put(300)
	fmt.Println(pool.Get().(int))

	if v, ok := pool.Get().(string); ok {
		fmt.Printf("value is %d,  convert ok....", v)
	} else {
		fmt.Printf("%T\n", v)
	}
}
