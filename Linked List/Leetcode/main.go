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

func main() {
	obj := Constructor()
	obj.AddAtHead(100)
	fmt.Println(obj)

}
