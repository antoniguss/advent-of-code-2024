package util

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func (list *LinkedList) Print() {
	fmt.Print("[")
	var current *Node = list.Head
	for current != nil {
		fmt.Printf("%d ", current.Value)
		current = current.Next
	}
	fmt.Println("]")

}
