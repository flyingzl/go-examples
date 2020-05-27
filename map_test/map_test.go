package map_test

import (
	"testing"
)

func TestMapInit(t *testing.T) {
	m1 := map[string]interface{}{"age": 10}
	t.Logf("len(m1) = %d", len(m1))
	if v, ok := m1["age"]; ok {
		t.Logf("m1[name]=%v", v)
	} else {
		t.Log("m1[2] doesn't exist")
	}
}

func TestMapTravel(t *testing.T) {
	m1 := map[string]interface{}{"age": 10, "name": "larry"}
	for key, val := range m1 {
		t.Logf("key = %v, value=%v", key, val)
	}
}
