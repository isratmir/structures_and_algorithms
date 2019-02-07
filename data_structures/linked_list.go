package main

import (
	"fmt"
	"sync"
)

type Node struct {
	Post string
	next *Node
}

type List struct {
	length int
	head   *Node
	tail   *Node
	lock   sync.Mutex
}

func (list *List) Add(node *Node) {
	list.lock.Lock()
	if list.length == 0 {
		list.head = node
		list.tail = node
	} else {
		curNode := list.tail
		curNode.next = node
		list.tail = node
	}

	list.length++
	list.lock.Unlock()
}

func (list *List) InsertAt(pos int, n *Node) error {
	if pos < 0 || pos > list.length {
		return fmt.Errorf("Index out of bounds")
	}

	if pos == 0 {
		n.next = list.head
		list.head = n
	} else {
		head := list.head
		var prevNode *Node
		for i := 1; i < pos; i++ {
			prevNode = head
			head = head.next
		}
		n.next = head
		prevNode.next = n
	}
	list.length++
	return nil
}

func (list *List) Remove(post string) {
	list.lock.Lock()
	curNode := list.head
	var prevNode *Node

	if curNode.Post == post {
		list.head = curNode.next
	} else {
		for curNode.Post != post {
			prevNode = curNode
			curNode = curNode.next
		}

		prevNode.next = curNode.next
	}
	list.length--
	list.lock.Unlock()
}

func (list *List) IsEmpty() bool {
	return list.length == 0
}

func main() {
	list := List{}
	node1 := Node{"Post 1", nil}
	node2 := Node{"Post 2", nil}
	node3 := Node{"Post 3", nil}
	node4 := Node{"Post 4", nil}
	node5 := Node{"Post 5", nil}

	fmt.Println(list.IsEmpty())

	list.Add(&node1)
	list.Add(&node2)
	list.Add(&node3)
	list.Add(&node4)
	list.Add(&node5)

	fmt.Println(list.IsEmpty())

	loopList(list.head)

	half := Node{"Post 2,5", nil}
	_ = list.InsertAt(3, &half)

	loopList(list.head)

	list.Remove("Post 2")

	loopList(list.head)
}

func loopList(head *Node) {
	fmt.Println("-----Start loop-----")
	for {
		fmt.Println(head)
		if head.next == nil {
			break
		}
		head = head.next
	}
	fmt.Println("-----End loop-----")
}
