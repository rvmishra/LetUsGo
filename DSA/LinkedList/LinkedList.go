package LinkedList

import "fmt"

type LinkedList struct {
	Val  int
	Next *LinkedList
}

func (head *LinkedList) Add(newNode *LinkedList) {
	if head == nil {
		return
	}

	for head.Next != nil {
		head = head.Next
	}

	head.Next = newNode
}

func (head *LinkedList) Print() {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

func (head *LinkedList) Length() int {
	length := 0
	for head != nil {
		length++
		head = head.Next
	}
	return length
}

func New(val int) *LinkedList {
	node := new(LinkedList)
	node.Val = val
	return node
}

func (list *LinkedList) HasCycle() bool {
	slow := list
	fast := list

	for fast.Next != nil && fast.Next.Next != nil {
		if fast.Next == slow || fast.Next.Next == slow {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}
