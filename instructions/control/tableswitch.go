package control

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type TABLE_SWITCH struct {
	defaultOffset		int32
	low					int32
	high				int32
	jumpOffsets			[]int32
}

/**
	tableswitch 指令操作码的后面后0～3个字节的padding，以保证
	defaultOffset在字节码中的地址是4的倍数
 */
func (self *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	self.defaultOffset = reader.ReadInt32()
	self.low = reader.ReadInt32()
	self.high = reader.ReadInt32()
	jumpOffsetsCount := self.high - self.low + 1
	self.jumpOffsets = reader.ReadInt32s(jumpOffsetsCount)
}

func (self *TABLE_SWITCH) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()

	index := stack.PopInt()
	var offset int
	if index < self.low || index > self.high {
		offset = int(self.defaultOffset)
	} else {
		offset = int(self.jumpOffsets[index - self.low])
	}

	base.Branch(frame, offset)
}
