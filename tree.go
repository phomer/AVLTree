/*
    Description: Simple AVLTree implementation

    TODO: Currently just an unbalanced tree for integers. Extend...
*/
package AVLTree

import (
    "fmt"
)

// Overall container
type Tree struct {
    root *Node
}

// Each node in the tree
type Node struct {
    value int
    balance int
    left *Node
    right *Node
}

// TODO: Just cheating on the testing for the moment.
func Start() {
    fmt.Printf("Testing\n")

    tree := new(Tree)
    tree.Insert(26)
    tree.Insert(45)
    tree.Print()

    fmt.Printf("The Original Allocation?\n")
    tree.Print()
}

// Insert an element in the tree
func (tree *Tree) Insert(value int) {
    if tree.root == nil {
	tree.root = &Node{value: value}
	tree.Print()
    } else {
	fmt.Printf("Second Node\n")
	tree.root.Insert(value)
    }
}

// Insert a node
func (node *Node) Insert(value int) {
    if value == node.value {
	return 

    } else if (value < node.value) {
	if node.left == nil {
	    node.left = &Node{value: value}
	} else {
	    node.left.Insert(value)
	}

    } else {
	if node.right == nil {
	    node.right = &Node{value: value}
	} else {
	    node.right.Insert(value)
	}
    }
}

// Print the tree
func (tree Tree) Print() {
    if tree.root != nil {
	fmt.Printf("Tree:\n")
	tree.root.Print(1)
    } else {
	fmt.Printf("Empty Tree\n")
    }
}

// Print the current node
func (node Node) Print(depth int) {
    if node.left != nil {
	node.left.Print(depth+1)
    }

    pad := ""
    for i := 0; i < depth; i++  {
	pad += "    "
    }
    fmt.Printf("%sValue: %v\n", pad, node.value)

    if node.right != nil {
	node.right.Print(depth+1)
    }
}

