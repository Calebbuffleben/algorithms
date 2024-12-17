package lists

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func ReverseLinkedList(head *Node) *Node {
	var prev *Node = nil
	current := head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}

func PrintLinkedList(head *Node) {
	current := head
	for current != nil {
		fmt.Printf("%d -> ", current.Value)
		current = current.Next
	}
	fmt.Println("nil")
}
