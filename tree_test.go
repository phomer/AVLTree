/*
    Description: Test the AVLTree implementation
*/

package AVLTree_test

import ( 
    "fmt"
    "testing"
    "github.com/phomer/AVLTree"
)

// Test the basic tree functionality
func TestStart(t *testing.T) {
    fmt.Printf("Building a tree\n")

    for _,array := range TableData() {
	tree := new(AVLTree.Tree)

	for _,value := range array {
	    if !tree.Exists(value) {
		tree.Insert(value)
		tree.Print()
	    }
	}
	tree.Print()
    }
}

// Internal list of table driven data
func TableData() ([][]int) {
    data := [][]int{
	[]int { 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23 },
	[]int { 23, 22, 21, 20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 3, 1 },
	[]int { 7, 5, 23, 18, 3, 3, 1, 19, 102, -3, 66, 18, 19, 23, 2, 77, 34, 8 ,24, 36 },
    }
    return data
}
