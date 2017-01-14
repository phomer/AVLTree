/*
    Description: Test the AVLTree implementation

*/

package AVLTree_test

import ( 
    "fmt"
    "testing"
    "github.com/phomer/AVLTree"
)

func TestStart(t *testing.T) {
    fmt.Printf("Building a tree\n")

    for _,array := range TableData() {
	tree := new(AVLTree.Tree)

	for _,value := range array {
	    tree.Insert(value)
	}
	tree.Print()
    }
}

func TableData() ([][]int) {
    data := [][]int{
	[]int { 1, 2, 3, 4, 5, 6, 7 },
	[]int { 7, 6, 5, 4, 3, 3, 1 },
	[]int { 7, 5, 23, 18, 3, 3, 1, 19 },
    }
    return data
}
