package linked_list

import (
	"fmt"
	"sync"
)

type Book struct {
	Title string
	next  *Book
}

type List struct {
	count int
	head  *Book
	tail  *Book
	lock  sync.RWMutex
}

func (list *List) IsEmpty() bool {
	list.lock.RLock()
	defer list.lock.RUnlock()
	return list.head == nil
}

func (list *List) Size() int {
	list.lock.RLock()
	defer list.lock.RUnlock()
	if list.head == nil {
		return 0
	}

	book := list.head
	size := 1
	for {
		if book.next == nil {
			break
		}
		book = book.next
		size++
	}
	return size
}

func (list *List) Append(title string) {
	list.lock.Lock()

	book := Book{title, nil}
	if list.head == nil {
		list.head = &book
		list.tail = &book
	} else {
		list.tail.next = &book
		list.tail = &book
	}
	list.count++
	list.lock.Unlock()
}

func (list *List) AppendAt(pos int, title string) error {
	if pos < 0 || pos > list.count {
		return fmt.Errorf("Index out of bounds")
	}

	return nil
}
