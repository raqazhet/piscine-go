package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	n := l
	if n == nil {
		return nil
	}

	n.Next = ListSort(n.Next)

	if n.Next != nil && n.Data > n.Next.Data {
		n = travel(n)
	}
	return n
}

func travel(l *NodeI) *NodeI {
	m := l
	d := l.Next
	res := d

	for d != nil && l.Data > d.Data {
		m = d
		d = d.Next
	}

	m.Next = l
	l.Next = d
	return res
}
