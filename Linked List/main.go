package main

import "fmt"

// Tipe data Node
type Node struct {
	data int
	link *Node
}

//Goal
//- Buat sebuah variabel untuk menyimpan data alamat ke head linkedlist
//- variabel tersebut akan menjadi sebuah identitas / base dari linkedlist
type LinkedList struct {
	// head adalah sebuah variabel yang menyimpan alamat head dari seubah linkedlist
	head *Node
	// Untuk track panjang linkedlist
	length int
}

//Buat Method untuk insert data di awal linked list
//Bayangkan sudah ada 1 element didalam linkedlist yang kita buat (for simplicity)
//second adalah temporary variable yang menyimpan alamat head sebelum dimasukkan node baru
//assign head variabel ke alamat node baru
//assign link pada head terbaru ke second (temporary variabel)
func (l *LinkedList) Prepend(n *Node) {
	second := l.head
	// Assign head ke alamat node yang ingin kita masukkan
	l.head = n
	// Jika linkedlist panjangnya == 0,maka node akan menunjuk ke node itu sendiri
	// Jika linkedlist panjangnya > 0,maka node akan menunjuk ke node pada head yang sebelumnya
	if l.length == 0 {
		l.head.link = nil
	} else {
		l.head.link = second
	}
	l.length++
}

func (l *LinkedList) Append(n *Node) {
	a := l.head

	for a.link != nil {
		a = a.link
	}
	a.link = n
}

func main() {
	myList := LinkedList{}

	node1 := &Node{data: 50}
	node2 := &Node{data: 51}
	node3 := &Node{data: 52}
	node4 := &Node{data: 100}
	myList.Prepend(node2)
	myList.Prepend(node3)
	myList.Append(node4)
	myList.Prepend(node1)

	a := myList.head
	for a.link != nil {
		a = a.link
	}
	fmt.Println(a.data)

}
