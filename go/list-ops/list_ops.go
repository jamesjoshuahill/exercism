package listops

type IntList []int

func (l IntList) Foldr(f func(item, accumulator int) int, initial int) int {
	a := initial
	for i := l.Length(); i > 0; i-- {
		a = f(l[i-1], a)
	}
	return a
}

func (l IntList) Foldl(f func(accumulator, item int) int, initial int) int {
	a := initial
	for _, e := range l {
		a = f(a, e)
	}
	return a
}

func (l IntList) Filter(f func(item int) bool) IntList {
	n := make(IntList, l.Length())
	var j int
	for _, e := range l {
		if f(e) {
			n[j] = e
			j++
		}
	}
	return n[:j]
}

func (l IntList) Length() int {
	var c int
	for range l {
		c++
	}
	return c
}

func (l IntList) Map(f func(item int) int) IntList {
	for i, e := range l {
		l[i] = f(e)
	}
	return l
}

func (l IntList) Reverse() IntList {
	length := l.Length()
	n := make(IntList, length)
	var j int
	for i := length; i > 0; i-- {
		n[j] = l[i-1]
		j++
	}
	return n
}

func (l IntList) Append(other IntList) IntList {
	return l.Concat([]IntList{other})
}

func (l IntList) Concat(lists []IntList) IntList {
	c := l.Length()
	for _, e := range lists {
		c += e.Length()
	}
	n := make(IntList, c)

	var j int
	for _, e := range l {
		n[j] = e
		j++
	}
	for _, other := range lists {
		for _, e := range other {
			n[j] = e
			j++
		}
	}
	return n
}
