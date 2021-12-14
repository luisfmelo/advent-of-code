package types

type ErrEmptyStack struct{}

func (e ErrEmptyStack) Error() string {
	return "empty_stack_error"
}

type RuneStack struct {
	stack []rune
}

func NewRuneStack() *RuneStack {
	return &RuneStack{stack: []rune{}}
}

func (s *RuneStack) Push(e rune) {
	s.stack = append(s.stack, e)
}

func (s *RuneStack) Pop() (rune, error) {
	if s.Size() == 0 {
		return '-', ErrEmptyStack{}
	}

	e := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]

	return e, nil
}

func (s *RuneStack) Size() int {
	return len(s.stack)
}
