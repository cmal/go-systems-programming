// tree with delete and search

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func traverse(t *Tree) {
	if t == nil {
		return
	}
	traverse(t.Left)
	fmt.Print(t.Value, " ")
	traverse(t.Right)
}

func create(n int) *Tree {
	var t *Tree
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10*n; i++ {
		temp := rand.Intn(n)
		t = insert(t, temp)
	}
	return t
}

func search(t *Tree, n int) bool {
	// returns whether a n can be found in the tree
	if t == nil {
		return false
	}

	if t.Value == n {
		return true
	}

	return search(t.Left, n) || search(t.Right, n)
}

func remove(t *Tree, n int) *Tree {
	// remove an element from the tree
	if t == nil {
		return nil
	}
	if t.Value == n {
		return tranverseInsert(t.Left, t.Right)
	}
	if t.Value < n {
		return tranverseInsert()
	}
}



func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v == t.Value {
		return t
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
		return t
	}
	t.Right = insert(t.Right, v)
	return t
}

func traverseInsert(t *Tree, tsource *Tree) *Tree {
	if t == nil {
		return tsource
	}

	if tsource == nil {
		return t
	}
	
	// tsource.Value should not equal to t.Value
	if tsource.Value < t.Value {
		t.Left = insert(t.Left, tsource.Value)
	}

	if tsource.Value > t.Value {
		t.Right = insert(t.Right, tsource.Value)
	}
	
	t = traverseInsert(t, tsource.Left)
	t = traverseInsert(t, tsource.Right)
	return t
}

func main() {
	tree := create(30)
	traverse(tree)
	fmt.Println()
	fmt.Println("The value of the root of the tree is", tree.Value)

	in := search(tree, 22)
	fmt.Println("22 is in tree? ", in)
}
