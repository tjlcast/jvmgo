package loads

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

/**
	加载指令从局部变量表中获取变量。然后推入操作数栈顶
	[33]
	aload 系列指令操作ref引用类型变量
	dload 系列操作double类型变量
	fload 系列操作float类型变量
	iload 系列操作int类型变量
	lload 系列操作long类型变量
	xaload 系列操作数组类型
 */

type ILOAD struct { base.Index8Instruction }

type ILOAD_0 struct { base.NoOperandInstruction }

type ILOAD_1 struct { base.NoOperandInstruction }

type ILOAD_2 struct { base.NoOperandInstruction }

type ILOAD_3 struct { base.NoOperandInstruction }

func _iload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}

func (self *ILOAD) Execute(frame *rtda.Frame) {
	_iload(frame, uint(self.Index))
}

func (self *ILOAD_0) Execute(frame *rtda.Frame) {
	_iload(frame, 0)
}

func (self *ILOAD_1) Execute(frame *rtda.Frame) {
	_iload(frame, 1)
}

func (self *ILOAD_2) Execute(frame *rtda.Frame) {
	_iload(frame, 2)
}

func (self *ILOAD_3) Execute(frame *rtda.Frame) {
	_iload(frame, 3)
}
