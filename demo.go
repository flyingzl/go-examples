package main

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	fmt.Println("hello")

	a := 3
	fmt.Println(a)

	{
		a := "d"
		fmt.Println(a)
	}

	b := map[int]string{1: "hello", 2: "love"}
	c := []string{"我x", "baby"}
	value, ok := interface{}(b).([]string)
	if ok {
		fmt.Println(value[1])
	} else {
		fmt.Println("===========")
		name := c[0]
		fmt.Println(len(name))                   // 4
		fmt.Println(len([]rune(name)))           // 2
		fmt.Println(strings.Count(name, "") - 1) // 2
		fmt.Println([]byte(name))
		fmt.Println(bytes.Count([]byte(name), nil) - 1) // 2
		fmt.Println(utf8.RuneCountInString(name))       // 2
		fmt.Println([]rune(name))

		type MyString = string
		e, ok := interface{}("hello").(MyString)
		fmt.Println(e, ok)
		fmt.Println(UnicodeIndex("你好呀l", "l"))
		fmt.Println((string)([]rune("你好")[0:1]))
	}

	whatAmi := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("int")
		case string:
			fmt.Println("string")
		case map[int]string:
			fmt.Println("map[int]string")
		default:
			fmt.Printf("Unkown type: %T\n", t)
		}
	}
	whatAmi(b)

	for _, i := range []int{3, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("failed:", e)
		} else {
			fmt.Println("ok", r)
		}
	}

	_, e := f2(42)
	// 准确的检测类型
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae)
	}

	if e != nil {
		fmt.Println(e)
	}
}

// UnicodeIndex 返回Unicode字符串中的索引
func UnicodeIndex(str, substr string) int {
	result := strings.Index(str, substr)
	if result >= 0 {
		result = utf8.RuneCountInString(str[0:result])
	}
	return result
}

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't be with 42")
	}
	return arg * 3, nil
}

type argError struct {
	arg int
	msg string
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't be with 42"}
	}
	return arg * 3, nil
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d-%s", e.arg, e.msg)
}
