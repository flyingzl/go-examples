package array_test

import (
	"testing"
)

func TestArrayInit(t *testing.T) {
	m1 := [...]int{1, 2, 3, 4, 5}
	m2 := [5]int{1, 2, 3, 4, 5}
	t.Log(m1 == m2)
	t.Logf("%T", m1)
}

func TestArrayTravel(t *testing.T) {
	m1 := [...]int{1, 2, 3, 4, 5}
	for index, val := range m1 {
		t.Logf("index = %v, value=%v", index, val)
	}
}
