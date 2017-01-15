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

    var left_balance = node.balance
    var balance = result.balance

    node.balance = left_balance - 1 - Max(balance, 0)
    result.balance = Min3(left_balance - 2, balance + left_balance - 2, balance - 1)

    fmt.Printf("Rotated Left to %v [%v->%v] from %v [%v->%v]\n", 
	result.value, balance, result.balance, 
	node.value, left_balance, node.balance)

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

    var right_balance = node.balance
    var balance = result.balance

    node.balance = right_balance + 1 - Min(balance, 0)
    result.balance = Max3(right_balance + 2, balance + right_balance + 2, balance + 1)

    fmt.Printf("Rotated Right to %v [%v->%v] from %v [%v->%v]\n", 
	result.value, balance, result.balance, 
	node.value, right_balance, node.balance)

    return result
}

// Insert a node
func (node *Node) Insert(value int) (*Node, int) {

    // Terminal Condition, create this node
    if node == nil {
	return &Node{value: value, balance: 0}, 1
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

    fmt.Printf("for value %v at value %v balance %v change %v\n", value, node.value, node.balance, change)

    node.balance += change

    // Rebalance at the parents
    var insert int = 0

    if node.balance != 0 && change != 0 {
	switch {

	case node.balance < -1:
	    node.Print(16)
	    if node.left.balance >= 0 {
		node.left = node.left.RotateLeft()
	    }
	    node = node.RotateRight()
	    node.Print(16)
	    insert = 0

	case node.balance > 1:
	    node.Print(16)
	    if node.right.balance <= 0 {
		node.right = node.right.RotateRight()
	    }
	    node = node.RotateLeft()
	    node.Print(16)
	    insert = 0

	default:
	    insert = 1
	}
    } else if change != 0 {
	insert = 0
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

// Print the current nodes, rotated 90 degrees, in-order traversal.
func (node Node) Print(depth int) {

    if node.right != nil {
	node.right.Print(depth+1)
    }

    padding := padding(depth)
    fmt.Printf("%sValue: %v [%v]\n", padding, node.value, node.balance)

    if node.left != nil {
	node.left.Print(depth+1)
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
func Max3(x int, y int, z int) (int) {
    return Max(Max(x,y),z)
}
