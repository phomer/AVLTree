/*
    Description: Simple AVLTree implementation
*/
package AVLTree

import (
    "fmt"
)

// Overall Tree Structure
type Tree struct {
    root *Node
}

// Insert an element in the tree
func (tree *Tree) Insert(value int) {
    tree.root, _ = tree.root.Insert(value)
}

// Test to see if this is in the tree or not. 
func (tree Tree) Exists(value int) (bool) {
    if tree.root == nil || tree.root.Find(value) == nil {
	return false;
    }
    return true;
}

// Remove, return true if successful
func (tree *Tree) Delete(value int) (bool) {
    if tree.root == nil {
	return false
    }
    // TODO: Wrong
    if tree.root.Delete(value) != tree.root {
	return true
    }

    return false
}

// Print the tree
func (tree Tree) Print() {
    if tree.root == nil {
	fmt.Printf("Empty Tree\n")
	return
    }

    fmt.Printf("--- AVL Tree:\n")
    tree.root.Print(1)
}

/*
    Tree Nodes
*/

// Each node in the tree
type Node struct {
    value int

    balance int
    left *Node
    right *Node
}

// Compare two values, 
func (node Node) Compare(value int) (int) {
    return node.value - value
}

// Walk the tree, searching
func (node *Node) Find(value int) (*Node) {
    if node == nil {
	return nil
    }

    if node.Compare(value) == 0 {
	return node;
    }

    var result *Node = nil

    if node.left != nil {
	result = node.left.Find(value)
    }

    if result == nil && node.right != nil {
	result = node.right.Find(value)
    }

    return result
}

// Rotate the tree to the left
func (node *Node) RotateLeft() (*Node) {
    if node == nil {
	return nil
    }

    if node.right == nil {
	return node
    }

    var result *Node = node.right
    node.right = result.left
    result.left = node

    var left_balance = result.balance
    var balance = node.balance

    node.balance = balance -1 - Max(left_balance,0)
    result.balance = Min3(balance-2, balance+left_balance-2, balance-1)

    return result
}

// Rotate the tree to the right
func (node *Node) RotateRight() (*Node) {
    if node == nil {
	return nil
    }

    if node.left == nil {
	return node
    }

    var result *Node = node.left
    node.left = result.right
    result.right = node

    var right_balance = result.balance
    var balance = node.balance

    node.balance = balance -1 - Max(right_balance,0)
    result.balance = Min3(balance-2, balance+right_balance-2, balance-1)

    return result
}

// Insert a node
func (node *Node) Insert(value int) (*Node, int) {

    // Terminal Condition, create this node
    if node == nil {
	return &Node{value: value}, 1
    }

    var change int = 0

    diff := node.Compare(value) 
    switch {

    case diff == 0:
	// Ignore duplicates

    case diff > 0:
    	node.left, change = node.left.Insert(value)
	change *= -1

    case diff < 0:
	node.right, change = node.right.Insert(value)
    }

    //fmt.Printf("for value %v balance %v change %v\n", value, node.balance, change)

    // Rebalance at the parent
    var insert int = 0
    var balance = node.balance + change

    if balance != 0 && change != 0 {
	switch {

	case balance < -1:
	    if node.left.balance < 0 {
		node.left = node.left.RotateLeft()
	    }
	    node = node.RotateRight()
	    insert = 0

	case balance > 1:
	    if node.right.balance > 0 {
		node.right = node.right.RotateRight()
	    }
	    node = node.RotateLeft()
	    insert = 0

	default:
	    insert = 1
	    node.balance = balance
	}
    }

    return node, insert
}

// Delete a node
func (node *Node) Delete(value int) (*Node) {
    if node == nil {
	return nil
    }

    if node.Compare(value) == 0 {
	if (node.left == nil) {
	    return node.right
	} else {
	    if (node.right == nil) {
		return node.left
	    }
	    return Merge(node.left, node.right)
	}
    }

    if node.left != nil {
	node.left = node.left.Delete(value)

    } else if node.right != nil {
	node.right = node.right.Delete(value)
    }
    return node
}

// TODO: Temp
func Merge(left *Node, right *Node) (*Node) {
    return left
}

// Print the current node
func (node Node) Print(depth int) {
    if node.left != nil {
	node.left.Print(depth+1)
    }

    padding := padding(depth)
    fmt.Printf("%sValue: %v\n", padding, node.value)

    if node.right != nil {
	node.right.Print(depth+1)
    }
}

// Convert depth into spaces, 4 per
func padding (size int) (string) {
    result := ""
    for i := 0; i < size; i++  {
	result += "    "
    }
    return result
}

func Max(x int, y int) (int) {
    if x > y {
	return x
    }
    return y
}

func Min(x int, y int) (int) {
    if x > y {
	return y
    }
    return x
}

func Min3(x int, y int, z int) (int) {
    return Min(Min(x,y),z)
}
