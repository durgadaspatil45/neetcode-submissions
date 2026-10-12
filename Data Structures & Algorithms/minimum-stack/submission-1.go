type MinStack struct {
	st    []int
	minSt []int
}

func Constructor() MinStack {
	return MinStack{
		st:    []int{},
		minSt: []int{},
	}
}

func (s *MinStack) Push(val int) {
	s.st = append(s.st, val)
	if len(s.minSt) == 0 || val <= s.minSt[len(s.minSt)-1] { // if it is empty and if the val is less than or equals to the current min top
		s.minSt = append(s.minSt, val)
	}
}

func (s *MinStack) Pop() {
	if len(s.st) == 0 {
		return
	}
	top := s.st[len(s.st)-1]
	s.st = s.st[:len(s.st)-1]

	if top == s.minSt[len(s.minSt)-1] {
		s.minSt = s.minSt[:len(s.minSt)-1]
	}
}

func (s *MinStack) Top() int {
	top := s.st[len(s.st)-1]
	return top
}

func (s *MinStack) GetMin() int {
	return s.minSt[len(s.minSt)-1]
}
