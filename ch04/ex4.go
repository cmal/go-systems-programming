// doubly linked list

package main

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}


func addNode(t *Node, v int) {
	if head == nil {
		t = &Node{v, nil, nil}
		head = t
		return
	}

	if v == t.Value {
		fmt.Println("Node already exists:", v)
		return
	}

	if t.Next == nil {
		t.Next = &Node{v, nil, t.Next}
		return
	}

	addNode(t.Next, v)
	return
}


func traverse(t *Node) {
	if t == nil {
		fmt.Println("-> Empty list!")
		return
	}

	for t != nil {
		fmt.Printf("<- %d -> ", t.Value)
		t = t.Next
	}
	fmt.Println()
}

var head = new(Node)

func main() {
	fmt.Println(head)  // global head is a pointer
	head = nil // global head is now a 0 pointer
	traverse(head) // head is nil
	addNode(head, 1)
	addNode(head, 1)
	traverse(head)
	addNode(head, 10)
	addNode(head, 5)
	addNode(head, 0)
	addNode(head, 0)
	traverse(head)
	addNode(head, 100)
	traverse(head)
}
