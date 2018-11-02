package extended

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

// branch if reference is null
type IFNULL struct {
	base.BranchInstruction
}

func (self *IFNULL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	ref := stack.PopRef()
	if ref == nil {
		base.Branch(frame, self.Offset)
	}
}

// branch if reference is not null
type IFNOTNULL struct {
	base.BranchInstruction
}

func (self *IFNOTNULL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	ref := stack.PopRef()
	if ref != nil {
		base.Branch(frame, self.Offset)
	}
}
