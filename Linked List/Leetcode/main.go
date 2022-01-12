package main

import "fmt"

type Node struct {
	data int
	next *Node
}
type singleList struct {
	head   *Node
	length int
}

func Constructor() *singleList {
	return &singleList{}
}

func (l *singleList) AddAtHead(val int) {
	node := Node{}
	node.data = val
	if l.head == nil {
		node.next = nil
	} else {
		node.next = l.head
	}
	l.head = &node
	l.length++
}

func main() {
	MyLinkedList := Constructor()
	MyLinkedList.AddAtHead(10)
	fmt.Println(MyLinkedList)

	MyLinkedList.AddAtHead(11)
	fmt.Println(MyLinkedList.head)
	// obj.AddAtHead(200)

}
