package extended

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

/**
	wide 指令增加了索引宽度，并不改变指令的操作，所以其
	Execute() 方法只要调用子指令的Execute() 方法即可
 */

type WIDE struct {
	modifiedInstruction base.Instruction
}

func (self *WIDE) FetchOperands(reader *base.BytecodeReader) {
	opcode := reader.ReadUint8()

	switch opcode {
	case 0x15: // iload
	case 0x16: // lload
	case 0x17: // fload
	case 0x18: // dload
	case 0x19: // aload
	case 0x36: // istore
	case 0x37: // lstore
	case 0x38: // fstore
	case 0x39: // dstore
	case 0x3a: // astore
	case 0x86: // iinc
	case 0xa9: // ret
		panic("Unsupported opcode: 0xs9!")
	}
}

func (self *WIDE) Execute(frame *rtda.Frame) {
	self.modifiedInstruction.Execute(frame)
}
