package piscine

func ListReverse(l *List) {
	if l.Head == nil {
		return
	}
	current := l.Head
	prev := l.Head
	prev = nil
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	temp := l.Head
	l.Head = l.Tail
	l.Tail = temp
}
