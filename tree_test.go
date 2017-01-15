/*
    Description: Test the AVLTree implementation
*/

package AVLTree_test

import ( 
    "fmt"
    "testing"
    "github.com/phomer/AVLTree"
)

// Test lots of inserts
func TestInsert(t *testing.T) {
    fmt.Printf("Creating a tree\n")

    for _,array := range TableData() {
	tree := new(AVLTree.Tree)

	for _,value := range array {
	    tree.Insert(value)
	}
	tree.Print()
    }
}

// test deleting the first tree
func TestDelete(t *testing.T) {
    fmt.Printf("Deleting a tree\n")

    list := TableData()

    tree := new(AVLTree.Tree)

    // Create the tree
    for _,value := range list[0] {
	tree.Insert(value)
    }
    tree.Print()

    // Take it apart in same order
    for _,value := range list[0] {
	status := tree.Delete(value)
	if status {
	    tree.Print()
	} else {
	    fmt.Printf("Not Found %v:\n", value)
	}
    }
} 


// Internal list of table driven data
func TableData() ([][]int) {
    data := [][]int{
	[]int { 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16,
	    17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31 },
	[]int { 23, 22, 21, 20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1 },
	[]int { 7, 5, 23, 18, 3, 3, 1, 19, 102, -3, 66, 18, 19, 23, 2, 77, 34, 8 ,24, 36, 20, 32, 17, 122, 4 },
    }
    return data
}
