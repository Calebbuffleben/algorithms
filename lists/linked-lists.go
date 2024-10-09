package lists

import ( 
	"fmt"
	"strings"
)

type node struct {
	value int
	next *node
}

type linkedList struct {
	head *node
	len int 
}

func (l *linkedList) add(value int) {
	newNode := new(node)
	newNode.value = value
	newNode.next = nil

	if l.head == nil {
		l.head = newNode
	} else {
		iterator := l.head

		for ; iterator.next != nil; iterator = iterator.next {

		}
		iterator.next = newNode
	}

}

func (l *linkedList) remove(value int) {

}

func (l linkedList) get(value int) *node {
	return nil
}

func (l *linkedList) String() string {
	sb := strings.Builder{}

	for iterator := l.head; iterator != nil; iterator = iterator.next {
		sb.WriteString(fmt.Sprintf("%d", iterator.value))
	}
	return sb.String()
}

func linkedLists() {
	l := linkedList{}
	l.add(1)
}