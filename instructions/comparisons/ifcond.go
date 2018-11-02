package comparisons

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

/**
	if<cond> 指令会把操作数栈顶的int变量弹出，然后跟0比较，
	满足条件进行跳转。
 */

type IFEQ struct {
	base.BranchInstruction
}

func (self *IFEQ) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()

	if val == 0 {
		base.Branch(frame, self.Offset)
	}
}

type IFNE struct {
	base.BranchInstruction
}

func (self *IFNE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()

	if val != 0 {
		base.Branch(frame, self.Offset)
	}
}

type IFLT struct {
	base.BranchInstruction
}

func (self *IFLT) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()

	if val < 0 {
		base.Branch(frame, self.Offset)
	}
}

type IFLE struct {
	base.BranchInstruction
}

func (self *IFLE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()

	if val <= 0 {
		base.Branch(frame, self.Offset)
	}
}

type IFGT struct {
	base.BranchInstruction
}

func (self *IFGT) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()

	if val > 0 {
		base.Branch(frame, self.Offset)
	}
}

type IFGE struct {
	base.BranchInstruction
}

func (self *IFGE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()

	if val >= 0 {
		base.Branch(frame, self.Offset)
	}
}
