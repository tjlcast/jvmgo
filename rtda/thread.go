package rtda

/**
	thread {
		pc 			uint32
		jvmStack	tstack
	}
 */

type Thread struct {
	pc			int
	stack		*Stack
}

/**
	new
 */
func NewThread() *Thread {
	return &Thread {
		stack: newStack(1024),
	}
}

/**
	get pc
 */
func (self *Thread) PC() int {
	return self.pc
}

/**
	set pc
 */
func (self *Thread) SetPC(pc int) {
	self.pc = pc
}

/**
	stack push
 */
func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}

/**
	stack pop
 */
func (self *Thread) PopFrame() *Frame {
	return self.stack.pop()
}

/**
	stack top
 */
func (self *Thread) CurrentFrame() *Frame {
	return self.stack.pop()
}

func (self *Thread) NewFrame(maxLocals, maxStack uint) *Frame {
	return newFrame(self, maxLocals, maxStack)
}
