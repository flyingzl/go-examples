package mutex_test

import (
	"sync"
	"testing"
)

func TestMutext(t *testing.T) {

	count := 0

	wg := new(sync.WaitGroup)

	wg.Add(2)

	mutext := new(sync.Mutex)

	go func() {
		for i := 0; i < 100000; i++ {
			mutext.Lock()
			count++
			mutext.Unlock()
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 100000; i++ {
			mutext.Lock()
			count--
			mutext.Unlock()
		}
		wg.Done()
	}()

	wg.Wait()
	t.Logf("now count=%d\n", count)

}
