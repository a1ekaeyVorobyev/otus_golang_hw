package list

import "fmt"

type Item struct {
	data interface{}
	next *Item
	prev *Item
	list *List
}

func (i *Item) Prev() *Item {
	return i.prev
}

func (i *Item) Next() *Item {
	return i.next
}

func (i *Item) GetList() *List {
	return i.list
}

func (i *Item) Value() interface{} {
	return i.data
}

type List struct {
	head *Item
	tail *Item
	size int
}

//Len() int // длина списка
func (l *List) Len() int {
	return l.size
}

// Remove(i *Item) // удалить элемент
func (l *List) Remove(remoteItem *Item) error {
	if remoteItem.list == nil {
		return fmt.Errorf("данный элемент  уже удален")
	}
	if l != remoteItem.list {
		return fmt.Errorf("данный элемент не принадлежит этому списку")
	}
	remoteItem.list = nil
	prev := remoteItem.prev
	next := remoteItem.next
	if prev != nil {
		prev.next = next
	} else {
		l.head = next
	}
	if next != nil {
		next.prev = prev
	} else {
		l.tail = prev
	}
	l.size--
	return nil
}

//PushBack(v interface{}) *Item // добавить значение в конец
func (l *List) PushBack(data interface{}) {
	if l == nil {
		l.New()
	}
	temp := l.Back()
	item := new(Item)
	item.data = data
	item.prev = temp
	l.tail = item
	if temp != nil {
		temp.next = item
	}
	if l.Front() == nil {
		l.head = item
	}
	item.list = l
	l.size++
}

//PushFront(v interface{}) *Item // добавить значение в начало
func (l *List) PushFont(data interface{}) {
	if l == nil {
		l.New()
	}
	temp := l.Front()
	item := new(Item) //:= new(Item)
	item.data = data
	item.prev = nil
	item.next = temp
	if temp != nil {
		temp.prev = item
	}
	l.head = item
	if l.Back() == nil {
		l.tail = item
	}
	item.list = l
	l.size++
}

//MoveToFront(i *Item) // переместить элемент в начало
func (l *List) MoveToFront(item *Item) {

	l.Remove(item)
	l.PushFont(item.data)
}

// Front() *Item // первый Item
func (l *List) Front() *Item {
	return l.head
}

// Back() *Item // последний Item
func (l *List) Back() *Item {
	return l.tail
}

func ListNew() *List {
	l := new(List)
	l.size = 0
	return l
}

func (l *List) New() {
	if l == nil {
		return
	}
	l = new(List)
	l.size = 0
	return
}

func (l *List) GetListItem(data interface{}) []*Item {
	temp := make([]*Item, 0, l.size)
	for e := l.head; e != nil; e = e.Next() {
		if e.data == data {
			temp = append(temp, e)
		}

	}
	return temp
}
