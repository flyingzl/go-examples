package unsafe_test

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"
)

func TestUnsafe(t *testing.T) {
	type MyInt int
	a := []int{1, 2, 3}
	t.Log(a)
	b := *(*[]MyInt)(unsafe.Pointer(&a))
	t.Log(b)
	a[0] = 2
	a[1] = 3
	t.Log(b)
}

func TestAtomic(t *testing.T) {
	buffer := new(unsafe.Pointer)
	var wg sync.WaitGroup

	writeFn := func() {
		data := []int{}
		for i := 0; i < 100; i++ {
			data = append(data, i)
		}
		atomic.StorePointer(buffer, unsafe.Pointer(&data))
	}

	readFn := func() {
		data := atomic.LoadPointer(buffer)
		fmt.Println(*(*[]int)(data))
	}

	writeFn()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			writeFn()
			time.Sleep(100 * time.Millisecond)
			wg.Done()
		}()
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			readFn()
			wg.Done()
		}()
	}

	wg.Wait()

}
