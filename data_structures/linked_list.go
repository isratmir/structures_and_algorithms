package main

import "fmt"

type Node struct {
	Post string
	next *Node
}

type List struct {
	length int
	head *Node
}

func (list *List) add(node *Node)  {

	if list.length == 0{
		list.head = node
	}else {
		curNode := list.head
		for curNode.next != nil{
			curNode = curNode.next
		}
		curNode.next = node
	}

	list.length++
}

func (list *List) remove(post string)  {
	curNode := list.head
	var prevNode *Node

	if curNode.Post == post {
		list.head = curNode.next
	} else {
		for curNode.Post != post{
			prevNode = curNode
			curNode = curNode.next
		}

		prevNode.next = curNode.next
	}
	list.length--
}

func main() {
	list := List{}
	node1 := Node{"Post 1", nil}
	node2 := Node{"Post 2", nil}
	node3 := Node{"Post 3", nil}
	list.add(&node1)
	list.add(&node2)
	list.add(&node3)

	n := list.head

	for n.next != nil{
		fmt.Println(n.next)
		n = n.next
	}
}