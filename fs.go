package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	err := ioutil.WriteFile("/tmp/data", []byte("你好hello"), 0644)
	check(err)

	v, err := ioutil.ReadFile("/tmp/data")
	check(err)
	fmt.Println(v)
	fmt.Println(string(v))

	f, err := os.OpenFile("/tmp/data", os.O_RDWR, 0644)
	defer f.Close()
	check(err)
	f.Seek(3, 0)
	b1 := make([]byte, 3)
	f.Read(b1)
	fmt.Println(string(b1))
	f.ReadAt(b1, 3)
	fmt.Println(string(b1))
	b1 = []byte("呀")
	f.WriteAt(b1, 6)
}
