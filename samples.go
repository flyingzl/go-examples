package main

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

// User info
type User struct {
	a int
	b bool
}

func (u User) toString() string {
	return fmt.Sprintf("%d(%t)", u.a, u.b)
}

func main() {

	showStruct()

	i := 0
	for i < 10 {
		i++
		defer deferFunc(i)
	}

	os := runtime.GOOS
	fmt.Println(os)
	today := time.Now().Weekday()

	switch time.Tuesday {
	case today:
		fmt.Println("today is Tuestday")
	case today + 1:
		fmt.Println("tomorrow")
	case today + 2:
		fmt.Println("the day after tomorrow")
	default:
		fmt.Println("too far away...")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("morning")
	case t.Hour() < 17:
		fmt.Println("afternoon")
	default:
		fmt.Println("evening...")
	}

	array := []int{1, 2, 3, 3}
	fmt.Printf("array type is %T, %v\n", array, array[1:2])
	// array = append(array, 4)
	fmt.Printf("array length is: %d\n", len(array))

	fmt.Println(strings.Fields("hello world"))
}

func deferFunc(i int) {
	fmt.Println("defer....", i)
}

func showStruct() {

	val := []User{
		{a: 3, b: false}, // 注意这里一定要有逗号
	}
	val = append(val, User{2, true})
	fmt.Println(val) // [{3 false} {2 true}]

	fmt.Println(User{1, true}.toString()) // 1(true)

	z := val[:]

	z[0].a = 4

	fmt.Println(val)

	for i, val := range z {
		fmt.Println(i, val)
		val.a = 3333
	}

	fmt.Println(z, val)

}
