package main

import (
	"LetUsGo/DSA/LinkedList"
	"fmt"
)

func main() {
	list := LinkedList.New(1)
	list.Add(LinkedList.New(2))
	list.Add(LinkedList.New(3))
	list.Add(LinkedList.New(3))
	list.Print()
	fmt.Println("length is ", list.Length())
	fmt.Println("Cycle Present : ", list.HasCycle())
}
