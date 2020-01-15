package main

func main() {
	// memoery list
	list := NewMemoryList(5)

	e1 := NewEntry("a", "b")
	e3 := NewEntry("e", "f")
	e4 := NewEntry("g", "h")
	list.PushFront(e1)
	list.PushFront(e3)
	list.PushFront(e4)
	list.Traverse()

	// Database list

	dblist := NewDBList(5)
	e2 := NewEntry("c", "d")
	e5 := NewEntry("u", "o")

	dblist.PushFront(e2)
	dblist.PushFront(e5)

	dblist.Traverse()
}
