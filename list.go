package main

import (
	"container/list"
	"fmt"
)

func main() {
	link := list.New()
	for i := 0; i < 5; i++ {
		link.PushBack(i)
	}

	link.PushFront(19)
	link.InsertAfter(20, link.Front())

	for el := link.Front(); el != nil; el = el.Next() {
		fmt.Println(el.Value)
	}

	link.Remove(link.Back())

	fmt.Printf("总共由%d个元素", link.Len())
}
