package rtda

type Stack struct {
	maxSize			uint
	size  			uint
	_top			*Frame
}

/**
	new stack
 */
func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}

/**
	push stack
 */
func (self *Stack) push(frame *Frame) {
	if self.size >= self.maxSize {
		panic("java.lang.StackOverflowError")
	}

	if self._top != nil {
		frame.lower = self._top
	}

	self._top = frame
	self.size++
}

/**
	pop stack
 */
func (self *Stack) pop() *Frame {
	if self._top == nil {
		panic("jvm stack is empty")
	}

	top := self._top
	self._top = top.lower
	top.lower = nil
	self.size--;
	return top
}

/**
	top stack
 */
func (self *Stack) top() *Frame {
	if self._top == nil {
		panic("jvm stack is empty")
	}
	return self._top
}
