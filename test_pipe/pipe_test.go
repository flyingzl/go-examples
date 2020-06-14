package test_pipe

import (
	filters "go-examples/test_pipe/filters"
	"testing"
)

func TestPipe(t *testing.T) {
	stringPipe := filters.NewStringPipe(",")
	toIntPipe := filters.NewToIntPipe()
	sumPipe := filters.NewSumPipe()
	pipeline := filters.NewStraightPipe("p1", stringPipe, toIntPipe, sumPipe)
	ret, err := pipeline.Process("1,2,3,4,5,6,7,8,9,10")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(ret.(int))
	}
}
