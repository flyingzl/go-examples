package mutex_test

import (
	"sync"
	"testing"
	"time"
)

func TestChanMessage(t *testing.T) {

	wg := new(sync.WaitGroup)
	c := make(chan int)

	wg.Add(1)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
		wg.Done()
	}()

	wg.Add(1)

	go func() {
		for v := range c {
			t.Log(v)
		}
		wg.Done()
	}()

	wg.Wait()

}

func TestChan(t *testing.T) {
	done := make(chan bool)

	go func() {
		time.Sleep(1 * time.Second)
		done <- true
	}()

	select {
	case <-done:
		t.Log("done...")
	case <-time.After(2 * time.Second):
		t.Log("time out..")
	}

}

func TestChannelCancel(t *testing.T) {

	c := make(chan bool)

	wg := new(sync.WaitGroup)

	isCanceled := func(c chan bool) bool {
		select {
		case <-c:
			return true
		default:
			return false
		}
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			for {
				if isCanceled(c) {
					t.Logf("%d is canceled", i)
					wg.Done()
					break
				} else {
					time.Sleep(1 * time.Second)
				}
			}
		}(i)
	}

	time.AfterFunc(time.Second*3, func() {
		close(c)
	})

	wg.Wait()

}
