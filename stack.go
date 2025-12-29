package main

type Stack struct {
	container []uint16
}

func (stack *Stack) Push(value uint16) {
	stack.container = append(stack.container, value)
}

func (stack *Stack) Pop() uint16 {
	stackLength := len(stack.container)
	return stack.container[stackLength-1]
}

func (stack *Stack) Peek() uint16 {
	//TODO: don't required for now
	return uint16(0x0000)
}

func (stack *Stack) IsEmpty() uint16 {
	//TODO: don't required for now
	return uint16(0x0000)
}
