package control

import (
	"jvmgo/rtda"
	"jvmgo/instructions/base"
)

/**
	无条件跳转
 */

type GOTO struct {
	base.BranchInstruction
}

func (self *GOTO) Execute(frame *rtda.Frame)  {
	base.Branch(frame, self.Offset)
}
