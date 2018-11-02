package constants

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type NOP struct {
	base.NoOperandInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {
	// 什么也不做
}