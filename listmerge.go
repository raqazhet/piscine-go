package piscine

func ListMerge(l1 *List, l2 *List) {
	if l1.Head == nil {
		l1.Head = l2.Head
		l1.Tail = l2.Tail
		return
	} else if l2.Head == nil {
		return
	}
	n := l1.Head
	for n.Next != nil {
		n = n.Next
	}
	n.Next = l2.Head
	l1.Tail = l2.Tail
}
