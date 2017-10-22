package utils

type entry struct {
	key   string
	value string
}

type List []entry

func (l List) Len() int {
	return len(l)
}

func (l List) Swap(i int, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l List) Less(i int, j int) bool {
	if l[i].value == l[j].value {
		return l[i].key < l[j].key
	} else {
		return l[i].value < l[j].value
	}
}
