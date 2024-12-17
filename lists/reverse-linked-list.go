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

func main() {
	// Create a sample linked list: 1 -> 2 -> 3 -> 4 -> nil
	head := &Node{Value: 1, Next: &Node{Value: 2, Next: &Node{Value: 3, Next: &Node{Value: 4, Next: nil}}}}

	fmt.Println("Original Linked List:")
	PrintLinkedList(head)

	// Reverse the linked list
	reversedHead := ReverseLinkedList(head)

	fmt.Println("Reversed Linked List:")
	PrintLinkedList(reversedHead)
}
