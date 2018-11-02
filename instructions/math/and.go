package math

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type IAND struct {
	base.NoOperandInstruction
}

func (self *IAND) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 & v2
	stack.PushInt(result)
}

type LAND struct {
	base.NoOperandInstruction
}

func (self *LAND) Execute(frame *rtda.Frame) {
	// todo
}
