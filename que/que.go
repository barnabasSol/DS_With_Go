package que

type Que struct {
	items []int
}

func (q *Que) Enqeue(item int) {
	q.items = append(q.items, item)
}

func (q *Que) Deqeue() int {
	deqd := q.items[0]
	q.items = q.items[1:]
	return deqd
}
