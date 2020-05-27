package slice_test

import (
	"testing"
)

func TestSliceInit(t *testing.T) {
	m1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	m2 := m1[:4]
	m3 := m1[3:]
	m2[3] = 0
	t.Logf("len(m2)=%d, cap(m2)=%d", len(m2), cap(m2))
	t.Logf("len(m3)=%d, cap(m3)=%d", len(m3), cap(m3))
}

func TestMapTravel(t *testing.T) {
	m1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for index, val := range m1 {
		t.Logf("key = %v, value=%v", index, val)
	}
}
