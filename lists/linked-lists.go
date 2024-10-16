package lists

import ( 
	"fmt"
	"strings"
)

type node struct {
	value int
	next *node
}

func (n node) String() string {
	return fmt.Sprintf("%d", n.value)
}

type linkedList struct {
	head *node
}

func (l *linkedList) add(value int, index int) {
	newNode := &node{value: value}

	if index == 0 {
		newNode.next = l.head
		l.head = newNode
		return
	}

	current := l.head
	counter := 0

	//traverse de linked list and assign the value of the next to the current
	for current != nil && counter < index - 1 {
		current = current.next
		counter++
	}

	if current == nil {
        fmt.Println("Index out of range")
        return
    }

	// Assign the new value to the current
	newNode.next = current.next
	current.next = newNode
}

func (l *linkedList) remove(value int) {
	var previous *node

	for current := l.head; current !=nil; current = current.next {
		if current.value == value {
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
		if iteractor.value == value {
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