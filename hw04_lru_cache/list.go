package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	len        int
	head, tail *ListItem
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.head
}

func (l *list) Back() *ListItem {
	return l.tail
}

func (l *list) PushFront(v interface{}) *ListItem {
	item := &ListItem{v, l.head, nil}
	if l.head != nil {
		l.head.Prev = item
	} else {
		l.tail = item
	}
	l.head = item
	l.len++
	return item
}

func (l *list) PushBack(v interface{}) *ListItem {
	item := &ListItem{v, nil, l.tail}
	if l.tail != nil {
		l.tail.Next = item
	} else {
		l.head = item
	}
	l.tail = item
	l.len++
	return item
}

func (l *list) Remove(i *ListItem) {
	switch {
	case l.head == l.tail:
		if i == l.head {
			l.head = nil
			l.tail = nil
		}
	case i == l.head:
		l.head = l.head.Next
		l.head.Prev = nil
	case i == l.tail:
		l.tail = l.tail.Prev
		l.tail.Next = nil
	default:
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
	}
	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	if i == l.head {
		return
	}
	l.Remove(i)
	l.len++
	l.head.Prev = i
	i.Next = l.head
	l.head = i
}

func NewList() List {
	return new(list)
}
