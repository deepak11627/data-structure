// Package singlyLinkedList implements a singly linked list
package main

import "fmt"

type DBNode struct {
	Value interface{}
	Next  *DBNode
}

// LinkedList sturct
type DBList struct {
	Database map[int]*DBNode
	NextKey  int
	Size     int
}

func NewDBList(size int) ListService {
	return &DBList{Size: size, Database: make(map[int]*DBNode, 0)}
}

func (l *DBList) Traverse() {

	for _, n := range l.Database {
		fmt.Println(n.Value)
	}
}

func (l *DBList) PushFront(i EntryService) {
	element := &DBNode{Value: i}
	// if len(l.Database) == 0 {
	l.NextKey++
	// 	l.Database[l.NextKey] = element
	// } else {
	// 	l.Database[l.NextKey].Next = element
	// }
	l.Database[l.NextKey] = element
	i.SetList(l)
}
