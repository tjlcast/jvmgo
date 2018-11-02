package comparisons

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

/**
	把栈顶的两个引用弹出，根据引用是否相同进行跳转。
 */

type IF_ACMPEQ struct {
	base.BranchInstruction
}

func (self *IF_ACMPEQ) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	rval1 := stack.PopRef()
	rval2 := stack.PopRef()

	if rval1 == rval2 {
		base.Branch(frame, self.Offset)
	}
}

type IF_ACMPNE struct {
	base.BranchInstruction
}

func (self *IF_ACMPNE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	rval2 := stack.PopRef()
	rval1 := stack.PopRef()

	if rval1 != rval2 {
		base.Branch(frame, self.Offset)
	}
}


