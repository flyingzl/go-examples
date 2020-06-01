package hello

import (
	"fmt"
)

func init() {
	fmt.Println("init....")
}

func init() {
	fmt.Println("init again....")
}

// Greet 用于显示招呼信息
func Greet(message string) {
	fmt.Printf("hello,  %s\n", message)
}
