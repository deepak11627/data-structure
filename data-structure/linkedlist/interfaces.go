package main

type NodeService interface {
	Next()
}

type ListService interface {
	PushFront(i EntryService)
	Traverse()
}

type EntryService interface {
	Next() EntryService
	Value() EntryService
	SetList(ListService)
}

type Entry struct {
	key  interface{}
	val  interface{}
	List ListService
}

func NewEntry(k, v interface{}) *Entry {
	return &Entry{key: k, val: v}
}

func (e *Entry) Next() EntryService {
	return nil
}

func (e *Entry) Value() EntryService {
	return nil
}

func (e *Entry) SetList(l ListService) {
	return
}
