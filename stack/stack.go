package stack

type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() int {
	len := len(s.items) - 1
	popped := s.items[len]
	s.items = s.items[:len]
	return popped
}
