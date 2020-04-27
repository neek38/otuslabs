package hw04_lru_cache //nolint:golint,stylecheck

type List interface {
	Len() int                          // длина списка
	Front() *listItem                  // первый Item
	Back() *listItem                   // последний Item
	PushFront(v interface{}) *listItem // добавить значение в начало
	PushBack(v interface{}) *listItem  // добавить значение в конец
	Remove(i *listItem)                // удалить элемент
	MoveToFront(i *listItem)           // переместить элемент в начало
}

type listItem struct {
	Value      interface{}
	next, prev *listItem
}

type list struct {
	first, last *listItem
	len         int
}

func (n *listItem) Next() *listItem {
	return n.next
}

func (n *listItem) Prev() *listItem {
	return n.prev
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *listItem {
	return l.first
}

func (l *list) Back() *listItem {
	return l.last
}

func (l *list) PushFront(v interface{}) *listItem {
	n := &listItem{Value: v}
	if l.first == nil {
		l.last = n
		l.len = 1
	} else {
		n.next = l.first
		l.first.prev = n
		l.len++
	}
	l.first = n
	return n
}

func (l *list) PushBack(v interface{}) *listItem {
	n := &listItem{Value: v}
	if l.last == nil {
		l.first = n
		l.len = 1
	} else {
		n.prev = l.last
		l.last.next = n
	}
	l.last = n
	l.len++
	return n
}

func (l *list) MoveToFront(i *listItem) {
	if i.prev != nil {
		i.prev.next = i.next
		if i.next != nil {
			i.next.prev = i.prev
		} else {
			l.last = i.prev
		}
		l.first.prev = i
		i.next = l.first
		i.prev = nil
		l.first = i
	}
}

func (l *list) Remove(i *listItem) {
	if l.len == 1 {
		l.first = nil
		l.last = nil
		l.len = 0
		return
	} else if i.prev == nil {
		i.next.prev = nil
		l.first = i.next
	} else if i.next == nil {
		i.prev.next = nil
		l.last = i.prev
	} else {
		i.prev.next = i.next
		i.next.prev = i.prev
	}
	l.len--
}

func NewList() List {
	return &list{}
}
