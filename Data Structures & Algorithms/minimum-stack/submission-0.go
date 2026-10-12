type MinStack struct {
	St []int
}

func Constructor() MinStack {
	stack := MinStack{}
	return stack
}

func (s *MinStack) Push(val int) {
	s.St = append(s.St, val)
}

func (s *MinStack) Pop() {
	s.St = s.St[:len(s.St)-1]
}

func (s *MinStack) Top() int {
	top := s.St[len(s.St)-1]
	return top
}

func (s *MinStack) GetMin() int {
	if len(s.St) == 0 {
		fmt.Println("The Stack is empty")
		return 0
	}
	min := s.St[0]
	for _, val := range s.St {
		if val < min {
			min = val
		}
	}
	return min
}