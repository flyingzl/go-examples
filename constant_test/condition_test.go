package constant_test

import (
	"testing"
)

const (
	Monday = iota + 1
	Tuesday
)

// 注意常量 定义没有 逗号， iota只可以在常量中 使用
const (
	Male = iota << 1
	Female
	Unknown
)

func TestConstant(t *testing.T) {
	t.Log(Monday, Tuesday)
	t.Log(Male, Female, Unknown)
	gendar := 2
	t.Log(gendar == Female)
}
