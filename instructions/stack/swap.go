package stack

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

/**
	主要对操作数栈的栈顶 slot 进行交换操作
 */

type SWAP struct {
	base.NoOperandInstruction
}

func (self *SWAP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}


