// Package singlyLinkedList implements a singly linked list
package main

import "fmt"

// Node struct
type Node struct {
	Next *Node

	Element interface{}
}

// LinkedList sturct
type MemoryList struct {
	Head *Node

	Size int
}

func NewMemoryList(size int) ListService {
	return &MemoryList{Size: size}
}
func (l *MemoryList) Traverse() {
	if l.Head == nil {
		return
	}
	n := *l.Head
	for {
		fmt.Println(n.Element)
		if n.Next == nil {
			return
		}
		n = *n.Next

	}
}

func (l *MemoryList) PushFront(i EntryService) {

	n := &Node{Element: i}
	if l.Head == nil {
		n.Next = nil
		l.Head = n
	} else {
		n.Next = l.Head
		l.Head = n
	}
	i.SetList(l)
}

func (l *MemoryList) Delete(i int) {
	// TODO
}

func (l *MemoryList) Search(i int) {
	// TODO
}
