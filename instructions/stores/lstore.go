package stores

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

/**
	存储指令把变量从操作数栈顶弹出，然后存入局部变量表中
 */


type LSTORE struct { base.Index8Instruction }

type LSTORE_0 struct { base.Index8Instruction }

type LSTORE_1 struct { base.Index8Instruction }

type LSTORE_2 struct { base.Index8Instruction }

type LSTORE_3 struct { base.Index8Instruction }

func _lstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index, val)
}

func (self *LSTORE) Execute(frame *rtda.Frame) {
	_lstore(frame, uint(self.Index))
}

//
//func (self *LSTORE_0) Execute(frame *rtda.Frame) {
//	_lstore(frame, uint(self.Index))
//}
//
//func (self *LSTORE_1) Execute(frame *rtda.Frame) {
//	_lstore(frame, uint(self.Index))
//}
//
//func (self *LSTORE_2) Execute(frame *rtda.Frame) {
//	_lstore(frame, uint(self.Index))
//}
//
//func (self *LSTORE_3) Execute(frame *rtda.Frame) {
//	_lstore(frame, uint(self.Index))
//}
