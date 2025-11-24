package Valid_Parentheses

type Stack struct {
	stack []rune
}

func (s *Stack) Push(v rune) {
	s.stack = append(s.stack, v)
}

func (s *Stack) Pop() rune {
	end := len(s.stack) - 1
	element := s.stack[end]
	s.stack = s.stack[:end]
	return element
}

func IsValid(s string) bool {
	stack := &Stack{}
	for _, ch := range s {
		if ch == '(' || ch == '[' || ch == '{' {
			stack.Push(ch)
		}
		if ch == ')' || ch == ']' || ch == '}' {
			if len(stack.stack) == 0 {
				return false
			}
			elem := stack.Pop()
			if elem == '(' && ch != ')' || elem == '[' && ch != ']' || elem == '{' && ch != '}' {
				return false
			}
		}
	}
	if len(stack.stack) > 0 {
		return false
	}

	return true
}
