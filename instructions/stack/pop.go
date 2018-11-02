package stack

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

/**
	栈指令直接对操作数栈进行操作，共9条：
		pop 和 pop2 把变量弹出
		dup 系列指令复制栈顶变量
		swap指令交换栈顶的两个变量
 */

type POP struct {
	base.NoOperandInstruction
}

type POP2 struct {
	base.NoOperandInstruction
}

func (self *POP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

func (self *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}