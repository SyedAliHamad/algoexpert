package main

import "fmt"

// Node represents a node in the BST.
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// Insert inserts a value into the BST.
func (n *Node) Insert(value int) *Node {
	if n == nil {
		return &Node{Value: value}
	}

	if value < n.Value {
		n.Left = n.Left.Insert(value)
	} else {
		n.Right = n.Right.Insert(value)
	}

	return n
}

// Find the closes value in the BST
func closestValueBST(value int, root *Node) int {

	for true {

		if value>root.Value{
			if root.Right==nil{
				return root.Value
			}
			root=root.Right
		} else if value<root.Value {
			if root.Left==nil{
				return root.Value
			}
			root=root.Left
		} else{
			return root.Value
		}
	}
	return 0
}

func main() {
	// Creating the BST with the specified values
	root := &Node{Value: 10}
	root.Insert(5)
	root.Insert(2)
	root.Insert(5)
	root.Insert(1)
	root.Insert(15)
	root.Insert(13)
	root.Insert(22)
	root.Insert(14)

	//Find Closest Value
	fmt.Println("Value closest to 25 is:", closestValueBST(25, root))
}
