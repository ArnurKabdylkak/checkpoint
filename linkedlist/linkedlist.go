package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}
type LinkedList struct {
	Head *ListNode
}

func (ll *LinkedList) Append(value int) {
	newNode := &ListNode{Val: value}
	if ll.Head == nil {
		newNode = ll.Head
		return
	}
	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}
	current = newNode
}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func (ll *LinkedList) Print() {
	current := ll.Head
	for current != nil {
		fmt.Printf("%d -> ", current.Val)
		current = current.Next
	}
	fmt.Println("nil")
}

func main() {
	ll := &LinkedList{} // Создаем пустой список

	ll.Append(1) // Добавляем элементы
	ll.Append(2)
	ll.Append(3)
	ll.Append(4)

	fmt.Print("Initial List: ")
	ll.Print()

	middle := middleNode(ll.Head)
	if middle != nil {
		fmt.Printf("Middle Node Value: %d\n", middle.Val)
	}
}
