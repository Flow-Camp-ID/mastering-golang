package utils // nama paket

import "fmt" // import paket fmt

// Mengimplementasikan linked list secara manual menggunakan struct dan pointer
func LinkedList() { // fungsi LinkedList
	fmt.Println("")
	fmt.Println("Linked List")
	fmt.Println("===========")

	// Node
	type Node struct { // deklarasi struct Node
		Name string // deklarasi variabel Name
		Next *Node  // deklarasi pointer Next
	}

	// Linked List
	var node4 *Node = &Node{Name: "Node 4", Next: nil}   // inisialisasi node4
	var node3 *Node = &Node{Name: "Node 3", Next: node4} // inisialisasi node3
	var node2 *Node = &Node{Name: "Node 2", Next: node3} // inisialisasi node2
	var node1 *Node = &Node{Name: "Node 1", Next: node2} // inisialisasi node1

	var currentNode *Node = node1 // inisialisasi currentNode

	for currentNode != nil { // perulangan selama currentNode tidak nil
		fmt.Println(currentNode.Name)  // print currentNode.Name
		currentNode = currentNode.Next // currentNode berpindah ke currentNode.Next
	}
}
