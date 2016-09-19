package main

import "fmt"

/*
	Struct Definition for binary Tree
*/
type Vertex struct {
	right, left *Vertex
	val int
}

/*
	Method to create Node for binary Tree
*/
func createVertex(val int) (*Vertex){
	return &Vertex{nil, nil, val}
}

/*
	PreOrder Traversal
*/
func InOrderTraversal(v *Vertex){
	if v == nil{
		return
	}
	fmt.Println(v.val)
	InOrderTraversal(v.left)
	InOrderTraversal(v.right)
}

/*
	Pre Order Traversal
*/
func PreOrderTraversal(v *Vertex) {
	if v == nil{
		return
	}
	PreOrderTraversal(v.left)
	fmt.Println(v.val)
	PreOrderTraversal(v.right)
}

/*
	Post Order Traversal
*/
func PostOrderTraversal(v *Vertex) {
	if v == nil{
		return
	}
	PostOrderTraversal(v.right)
	fmt.Println(v.val)
	PostOrderTraversal(v.left)
}

func main() {
	root := createVertex(2)
	left := createVertex(1)
	right := createVertex(3)
	root.left = left
	root.right = right
	fmt.Println(" In Order Traversal")
	InOrderTraversal(root)
	fmt.Println(" Pre Order Traversal")
	PreOrderTraversal(root)
	fmt.Println(" Post Order Traversal")
	PostOrderTraversal(root)
}

