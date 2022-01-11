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

func main() {
	// Inisialisais Linked List
	//Node ke 0,pada saat ini head dari myList = nil
	myList := LinkedList{}
	//Membuat Node
	node1 := &Node{data: 50}
	node2 := &Node{data: 51}
	node3 := &Node{data: 52}

	//Node ke 1,pada saat ini head dari myList = alamat dari node1 dan node1 link ke nil
	myList.Prepend(node1)
	//Node ke 2,pada saat ini head dari myList = alamat dari node2 dan node2 refrensi ke node1
	myList.Prepend(node2)
	//Node ke 3,pada saat ini head dari myList = alamat dari node3 dan node3 refrensi ke node2
	myList.Prepend(node3)

	//Akan print alamat yang sama
	// fmt.Println(node3)
	fmt.Println(myList.head.data)
}
