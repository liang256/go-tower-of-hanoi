package stack

type Stack struct {
	list []int
}

func New() *Stack {
	s := new(Stack)
	s.list = make([]int, 0, 8)
	return s
}

func (s *Stack) IsEmpty() bool {
	return len(s.list) == 0
}

func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	poped := s.list[len(s.list)-1]
	s.list = s.list[:len(s.list)-1]
	return poped, true
}

func (s *Stack) Push(element int) {
	s.list = append(s.list, element)
}

func (s *Stack) ToSlice() []int {
	return s.list
}

func (s *Stack) Size() int {
	return len(s.list)
}
