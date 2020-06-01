package duck_type_test

import (
	"errors"
	"fmt"
	"go-examples/hello"
	"testing"
)

func checkType(v interface{}) string {
	switch t := v.(type) {
	case string:
		return "string"
	case int:
		return "int"
	case bool:
		return "bool"
	default:
		return fmt.Sprintf("Unknown: %T", t)
	}
}

func conv(val interface{}) (string, error) {
	if v, ok := val.(string); ok {
		return v, nil
	} else {
		return "", errors.New(fmt.Sprintf("can't convert type %T to string", val))
	}
}

func TestInterfaceConv(t *testing.T) {
	t.Log(checkType(1))
	t.Log(checkType(false))
	t.Log(checkType("2"))
	t.Log(checkType([]int{}))
	if v, err := conv(1); err == nil {
		t.Log(v)
	} else {
		t.Error(err)
	}
}

func TestDefer(t *testing.T) {
	defer func() {
		t.Log("defer func invoked...")
		if err := recover(); err != nil {
			t.Error("found error: ", err)
		}
	}()
	t.Log("testing defer...")
	panic("not permited")
}
