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

func plus(num int) func() int {
	v := num
	return func() int {
		v++
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

func TestClosureFn(t *testing.T) {
	fn := plus(2)
	t.Log(fn())
	t.Log(fn())
	t.Log(fn())
	t.Log(concatString("x2", "3"))
}

func BenchmarkConcatString(b *testing.B) {
	s := []string{"a", "b", "c", "d", "e", "1", "2", "3", "4"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		concatString(s...)
	}
	b.StopTimer()
}

func BenchmarkConcatStringWithFor(b *testing.B) {
	s := []string{"a", "b", "c", "d", "e", "1", "2", "3", "4"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		concatStringWithFor(s...)
	}
	b.StopTimer()
}

func BenchmarkConcatStringWithBuffer(b *testing.B) {
	s := []string{"a", "b", "c", "d", "e", "1", "2", "3", "4"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		concatStringWithBuffer(s...)
	}
	b.StopTimer()
}

func BenchmarkConcatStringWithNew(b *testing.B) {
	s := []string{"a", "b", "c", "d", "e", "1", "2", "3", "4"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		concatStringWithNew(s...)
	}
	b.StopTimer()
}

func TestMultipleParams(t *testing.T) {
	sum := func(v ...int) int {
		ret := 0
		for _, value := range v {
			ret += value
		}
		return ret
	}

	a := []int{1, 3, 4}

	v := sum(a...)
	t.Log(v)
}
