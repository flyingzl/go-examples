package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	result := sum(1, 2)
	if result == 3 {
		t.Logf("the expected sum is %q", result)
	} else {
		t.Errorf("the extected sum shoule be 3, but the actual num is %q ", result)
	}
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum(i, 2)
	}
}
