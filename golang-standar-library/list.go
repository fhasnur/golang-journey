package main

import (
	"container/list"
	"fmt"
)

func main() {
	var data *list.List = list.New()

	data.PushBack("Fandi")
	data.PushBack("Meylwan")
	data.PushBack("Hasnur")

	var head *list.Element = data.Front()
	fmt.Println(head.Value) // Fandi

	next := head.Next()
	fmt.Println(next.Value) // Meylwan

	next = next.Next()
	fmt.Println(next.Value) // Hasnur

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
