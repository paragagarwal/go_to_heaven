package main

import "fmt"

/*
	Structure defintoion for Linked List
*/
type Node struct {
	value int
	next *Node
}

/*
	Create Node for linked list
*/
func createNode(value int) (*Node) {
	return &Node{value, nil}
}

/*
	Add node to linked list
*/
func addNode(value int, head *Node){
	pnode := head
	for ; pnode.next != nil; pnode = pnode.next {
	}
	pnode.next = createNode(value)
}

/*
	Add node to linked list such that elements are sorted
*/
func addSortNode(value int, head *Node){
	pnode := head
	for ; pnode.next != nil && pnode.next.value < value ; pnode = pnode.next {
	}
	tempNode := pnode.next
	pnode.next = createNode(value)
	pnode.next.next = tempNode
}

/*
	Method to set nill
*/
func setNil() (*Node){
	return nil
}

/*
	Reverse the Linked List
*/
func reverse(head *Node) (*Node) {
	if (head == nil || head.next == nil){
		return head
	}
	prevNode := setNil()
	nextNode := head
	for ; nextNode != nil ;  {
		tempNode := nextNode.next
		nextNode.next = prevNode
		prevNode = nextNode
		nextNode = tempNode
	}
	return prevNode
}

/*
	print the list in sequence
*/
func printList(head *Node, str string) {
	fmt.Println(str)
	for e := head; e != nil; e = e.next {
		fmt.Println(e.value)
	}
}

func main() {
	head := createNode(1)
	addSortNode(4, head)
	addSortNode(3, head)
	addSortNode(2, head)
	printList(head, "initial list")
	head = reverse(head)
	printList(head, "reverse list")
	}