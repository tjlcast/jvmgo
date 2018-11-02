package constants

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

/**
	bipush: 	指令从操作数中获取一个byte型整数，扩展成int型，然后推入栈顶。
	sipush:		指令从操作数中获取一个short型整数，扩展成int型，然后推入栈顶，
 */

type BIPUSH struct {
	val int8
}

type SIPUSH struct {
	val int16
}

func (self *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt8()
}

func (self *BIPUSH) Execute(frame *rtda.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}


