package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	temp := l.Head
	prv := l.Head
	for temp != nil && temp.Data == data_ref {
		l.Head = temp.Next
		temp = l.Head
	}
	for temp != nil {
		for temp != nil && temp.Data != data_ref {
			prv = temp
			temp = temp.Next
		}
		if temp == nil {
			return
		}
		prv.Next = temp.Next
		temp = prv.Next
	}
}
