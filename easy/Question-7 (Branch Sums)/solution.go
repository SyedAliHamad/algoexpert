package main

import "log"

func branchSums(root *Node) []int{
	var sums []int
	calculatebranchSum(root,0,&sums)
	return sums
}

func calculatebranchSum(node *Node, runningSum int, sums *[]int) {
	if node==nil{
		return
	}

	newRunningSum:=runningSum+node.Value
	if node.Left==nil && node.Right==nil {
		*sums = append(*sums, newRunningSum)
		log.Println(sums)
		return
	}

	calculatebranchSum(node.Left,newRunningSum,sums)
	calculatebranchSum(node.Right,newRunningSum,sums)
}

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
	arr:=branchSums(root)
	log.Println(arr)
}
