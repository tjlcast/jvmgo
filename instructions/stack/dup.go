package stack

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

/**


 */


type DUP struct {
	base.NoOperandInstruction
}

type DUP_X1 struct {
	base.NoOperandInstruction
}

type DUP_X2 struct {
	base.NoOperandInstruction
}

type DUP2 struct {
	base.NoOperandInstruction
}

type DUP2_X1 struct {
	base.NoOperandInstruction
}

type DUP2_X2 struct {
	base.NoOperandInstruction
}

func (self *DUP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot := stack.PopSlot()
	stack.PushSlot(slot)
	stack.PushSlot(slot)
}


func (self *DUP_X1) Execute(frame *rtda.Frame) {
	// do nothing
}

func (self *DUP_X2) Execute(frame *rtda.Frame) {
	// do nothing
}

func (self *DUP2) Execute(frame *rtda.Frame) {
	// do nothing
}

func (self *DUP2_X1) Execute(frame *rtda.Frame) {
	// do nothing
}

func (self *DUP2_X2) Execute(frame *rtda.Frame) {
	// do nothing
}
