package list

import "fmt"

type List struct {
	front *Node
	back  *Node
	size  int
}

func New() *List {
	return &List{}
}

func (l *List) InsertFront(value int) {
	newFront := &Node{value: value}
	if l.front == nil && l.back == nil {
		l.insertRoot(newFront)
		return
	}
	oldFront := l.front
	newFront.next = oldFront
	l.front = newFront
	oldFront.prev = newFront
	l.size++
}

func (l *List) InsertBack(value int) {
	newBack := &Node{value: value}
	if l.front == nil && l.back == nil {
		l.insertRoot(newBack)
		return
	}
	oldBack := l.back
	newBack.prev = oldBack
	l.back = newBack
	oldBack.next = newBack
	l.size++
}

func (l *List) insertRoot(node *Node) {
	l.front = node
	l.back = node
	l.size++
}

func (l *List) InsertAfterFromFront(valueAfter int, value int) error {
	n := l.front
	for n.value != valueAfter {
		if n.next == nil {
			return fmt.Errorf("valueAfter does not exist in the list")
		}
		n = n.next
	}
	if n == l.back {
		l.InsertBack(value)
		return nil
	}
	newNode := &Node{value: value}
	next := n.next
	newNode.prev = n
	newNode.next = next

	n.next = newNode
	if next != nil {
		next.prev = newNode
	}
	l.size++
	return nil
}

func (l *List) InsertBeforeFromBack(valueBefore int, value int) error {
	n := l.back
	for n.value != valueBefore {
		if n.prev == nil {
			return fmt.Errorf("valueBefore does not exist in the list")
		}
		n = n.prev
	}
	if n == l.front {
		l.InsertFront(value)
		return nil
	}
	newNode := &Node{value: value}
	prev := n.prev
	newNode.prev = prev
	newNode.next = n

	n.prev = newNode
	if prev != nil {
		prev.next = newNode
	}
	l.size++
	return nil
}

func (l *List) DeleteFromFront(value int) error {
	if len := l.Len(); len <= 1 {
		if len == 1 {
			l.deleteRoot()
			return nil
		} else {
			return fmt.Errorf("the list is empty")
		}
	}
	n := l.front
	if n.value == value {
		return l.DeleteFront()
	}
	// так как удаляем с начала, то правильно будет найти первый элемент
	// с совпадающим value, а только потом делать проверку на back
	for n.value != value {
		if n.next == nil {
			return fmt.Errorf("value does not exist in the list")
		}
		n = n.next
	}
	if n == l.back {
		return l.DeleteBack()
	}

	next := n.next
	prev := n.prev
	next.prev = prev
	prev.next = next
	l.size--
	return nil
}

func (l *List) DeleteFromBack(value int) error {
	if len := l.Len(); len <= 1 {
		if len == 1 {
			l.deleteRoot()
			return nil
		} else {
			return fmt.Errorf("the list is empty")
		}
	}
	n := l.back
	if n.value == value {
		return l.DeleteBack()
	}
	// так как удаляем с конца, то правильно будет найти последний элемент
	// с совпадающим value, а только потом делать проверку на front
	for n.value != value {
		if n.prev == nil {
			return fmt.Errorf("value does not exist in the list")
		}
		n = n.prev
	}
	if n == l.front {
		return l.DeleteFront()
	}

	next := n.next
	prev := n.prev
	next.prev = prev
	prev.next = next
	l.size--
	return nil
}

func (l *List) DeleteFront() error {
	if len := l.Len(); len <= 1 {
		if len == 1 {
			l.deleteRoot()
			return nil
		} else {
			return fmt.Errorf("the list is empty")
		}
	}
	n := l.front
	next := n.next
	next.prev = nil
	l.front = next
	l.size--
	return nil
}

func (l *List) DeleteBack() error {
	if len := l.Len(); len <= 1 {
		if len == 1 {
			l.deleteRoot()
			return nil
		} else {
			return fmt.Errorf("the list is empty")
		}
	}
	n := l.back
	prev := n.prev
	prev.next = nil
	l.back = prev
	l.size--
	return nil
}

func (l *List) deleteRoot() {
	l.front = nil
	l.back = nil
	l.size--
}

func (l *List) Front() (int, error) {
	if l.Len() <= 0 {
		return 0, fmt.Errorf("the list is empty")
	}
	return l.front.value, nil
}

func (l *List) Back() (int, error) {
	if l.Len() <= 0 {
		return 0, fmt.Errorf("the list is empty")
	}
	return l.back.value, nil
}

func (l *List) Len() int {
	return l.size
}

func (l *List) ToSlice() []int {
	result := make([]int, l.size)
	n := l.front
	for i := 0; i < l.size && n != nil; i++ {
		result[i] = n.value
		n = n.next
	}
	return result
}

func (l *List) ToSliceReverse() []int {
	result := make([]int, l.size)
	n := l.back
	for i := 0; i < l.size && n != nil; i++ {
		result[i] = n.value
		n = n.prev
	}
	return result
}
