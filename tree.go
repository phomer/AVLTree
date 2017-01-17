/*
    Description: Simple AVL Tree implementation
*/
package AVLTree

import (
    "fmt"
    "strconv"
)

// Overall Tree Structure
type Tree struct {
    root *Node
}

// Test to see if this is in the tree or not. 
func (tree Tree) Exists(value int) (bool) {
    if tree.root == nil || tree.root.Find(value) == nil {
	return false;
    }
    return true;
}

// Insert an element in the tree
func (tree *Tree) Insert(value int) {
    tree.root, _ = tree.root.Insert(value)
}

// Remove, return true if successful
func (tree *Tree) Delete(value int) (bool) {
    var change int = 0

    if tree.root == nil {
	return false
    }

    tree.root, change = tree.root.Delete(value)

    if change == 0 {
	return false
    }

    return true
}

// Update, not sure if this is what was intended...
func (tree *Tree) Update(orig int, value int) {
    tree.Delete(orig)
    tree.Insert(value)
}

// Print the tree
func (tree Tree) Print() {
    if tree.root == nil {
	fmt.Printf("--- AVL Tree:\n    EMPTY\n")
	return
    }

    fmt.Printf("--- AVL Tree:\n")
    tree.root.Print(1)
}

/*
    AVL Tree Nodes
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
    result.balance = Min(left_balance - 2, balance + left_balance - 2, balance - 1)

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
    result.balance = Max(right_balance + 2, balance + right_balance + 2, balance + 1)

    return result
}

// Insert a node
func (node *Node) Insert(value int) (*Node, int) {

    // Terminal Condition, create this node
    if node == nil {
	return &Node{value: value, balance: 0}, 1
    }

    var change int = 0

    // Descend to the children
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

    node.balance += change

    // Rebalance at the parents or grandparents
    var insert int = 0

    if node.balance != 0 && change != 0 {
	switch {

	case node.balance < -1:
	    if node.left.balance >= 0 {
		node.left = node.left.RotateLeft()
	    }
	    node = node.RotateRight()
	    insert = 0

	case node.balance > 1:
	    if node.right.balance <= 0 {
		node.right = node.right.RotateRight()
	    }
	    node = node.RotateLeft()
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
func (node *Node) Delete(value int) (*Node, int) {
    var change int = 0

    if node == nil {
	return nil, change
    }

    diff := node.Compare(value) 
    switch {
    case diff > 0:
	node.left, change = node.left.Delete(value)

    case diff < 0:
	node.right, change = node.right.Delete(value)
	change *= -1
	    
    case diff == 0:
	switch {
	case node.left == nil:
	    return node.right, 1

	case node.right == nil:
	    return node.left, 1

	default:
	    // Pick the heavier of the two...
	    if -1 * node.left.balance < node.right.balance {
		node = node.RotateLeft()
		node.left, change = node.left.Delete(value)

	    } else {
		node = node.RotateRight()
		node.right, change = node.right.Delete(value)
		change *= -1
	    }
	}
    }

    // Update the balance
    if change != 0 {

	if node.balance != change {
	    node.balance += change
	}

	switch {
	case node.balance < -1:
	    if node.left.balance >= 0 {
		node.left = node.left.RotateLeft()
	    }
	    node = node.RotateRight()

	case node.balance > 1:
	    if node.right.balance <= 0 {
		node.right = node.right.RotateRight()
	    }
	    node = node.RotateLeft()
	}
    }

    return node, change
}

// Print the current nodes, rotated 90 degrees, in-order traversal.
func (node Node) Print(depth int) {

    if node.right != nil {
	node.right.Print(depth+1)
    }

    padding := padding(depth)
    fmt.Printf("%s[%v] Value: %v\n", padding, symbols(node.balance), node.value)

    if node.left != nil {
	node.left.Print(depth+1)
    }
}

// Slightly more decorative: shift 1/-1 -> +/-
func symbols(balance int) (string) {
    switch {
    case balance == 1:
	return "+"
    case balance == -1:
	return "-"
    }
    return strconv.Itoa(balance)
}

// Convert depth into spaces, 4 per
func padding (size int) (string) {
    result := ""
    for i := 0; i < size; i++  {
	result += "    "
    }
    return result
}

func Max(values ...int) (int) {
    var total int = values[0]
    for _,value := range values {
	if (value > total) {
	    total = value
	}
    }
    return total
}

func Min(values ...int) (int) {
    var total int = values[0]
    for _,value := range values {
	if (value < total) {
	    total = value
	}
    }
    return total
}

