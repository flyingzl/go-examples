package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type IntConv func(value int) int

func calCost(fn IntConv) IntConv {
	start := time.Now()
	return func(n int) int {
		v := fn(n)
		fmt.Println(time.Since(start).Milliseconds())
		return v
	}
}

func TestSum(t *testing.T) {
	start := time.Now()
	result := sum(1, 2)
	t.Log(time.Since(start).Milliseconds())
	if result == 3 {
		t.Logf("the expected sum is %v", result)
	} else {
		t.Errorf("the extected sum shoule be 3, but the actual num is %q ", result)
	}
}

func TestVarParams(t *testing.T) {
	fn := func(a int) int {
		time.Sleep(time.Millisecond * time.Duration(rand.Int31n(100)))
		return a * a
	}

	newFn := calCost(fn)
	t.Logf("result is  %d", newFn(2))
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum(i, 2)
	}
}
