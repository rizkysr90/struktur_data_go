package main

import "fmt"

type MyLinkedList struct {
	data int
	link *MyLinkedList
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}
func (this *MyLinkedList) AddAtHead(val int) {
	node := Constructor()
	node = *this

	this.data = val
	if node.link == nil {
		this.link = nil
	} else {
		this.link = &node
	}
}
func (this *MyLinkedList) AddAtTail(val int) {
	node := Constructor()

	currNode := this
	for currNode.link != nil {
		currNode = currNode.link
	}
	currNode.link = &node
	node.data = val
	node.link = nil
}

func main() {
	obj := Constructor()
	obj.AddAtHead(100)
	obj.AddAtTail(200)
	fmt.Println(obj.link)

}
