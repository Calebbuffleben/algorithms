package lists

import ( 
	"fmt"
	"strings"
	"errors"
)

type node struct {
	data int
	next *node
}

func (n node) String() string {
	return fmt.Sprintf("%d", n.data)
}

type linkedList struct {
	head *node
	size int 
}

// NewLinkedList creates and returns a new empty linked list
func NewLinkedList() *linkedList {
    return &linkedList{head: nil, size: 0}
}

func (l *linkedList) add(position int, data int) error {
	if position < 0 || position > l.size {
		return errors.New("invalid position")
	}

	newNode := &node{data: data}

	if position == 0 {
		//Insert at beggining
		newNode.next = l.head
		l.head = newNode;
	} else {
		// Insert at the specified position
		current := l.head
		for i := 0; i < position - 1; i++{
			current = current.next
		}
		newNode.next = current.next
		current.next = newNode
	}

	l.size++
	return nil
}

func (l *linkedList) remove(value int) {
	var previous *node

	for current := l.head; current !=nil; current = current.next {
		if current.data == value {
			if l.head == current {
				l.head = current.next
			} else {
				previous.next = current.next
				return
			}
		}
		previous = current
	}
}

func (l linkedList) get(value int) *node {
	for iteractor := l.head; iteractor != nil; iteractor = iteractor.next {
		if iteractor.data == value {
			return iteractor
		}
	}

	return nil
}

func (l *linkedList) String() string {
	sb := strings.Builder{}

	for iterator := l.head; iterator != nil; iterator = iterator.next {
		sb.WriteString(fmt.Sprintf("%d", iterator.data))
	}
	return sb.String()
}

func linkedLists() {
	l := linkedList{}
	fmt.Println(l)
	fmt.Println(*(l.get(5)))
	l.remove(3)
	l.remove(5)
	fmt.Println(l)
}