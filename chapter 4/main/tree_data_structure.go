package main

import "fmt"

type Tree struct {
	value int
	left,right *Tree
}

type address struct {
	hostname string
	port int }


func main() {
	values := []int{13,35,27,88,14,3,7}
	var tree *Tree

	for _,v := range values {
		tree = add(tree,v)
	}

	tree_list(tree,values[:0])
	fmt.Println(values)

	///////////////////

	hits := make(map[address]int)
	hits[address{"golang.org", 443}]++

	fmt.Println(hits)
}


func add(t *Tree, value int) *Tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(Tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func tree_list(tree *Tree,values []int) []int {

	if tree != nil {
		tree_list(tree.left,values)
		//fmt.Println(tree.value)
		values = append(values,tree.value)
		tree_list(tree.right,values)
	}

	return values
}


