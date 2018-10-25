package base

import "jvmgo/rtda"

type Instruction interface {
	/**
		提取操作数
	 */
	 FetchOperands(reader *BytecodeReader)

	 /**
	 	执行指令逻辑
	  */
	 Execute(frame *rtda.Frame)
}

/**
	NoOperandInstruction
 */
type NoOperandInstruction struct {}

func (self *NoOperandInstruction) FetchOperands(reader *BytecodeReader) {
	// nothing to do
}

/**
	BranchInstruction
	这是一个跳转指令
 */
type BranchInstruction struct {
	Offset			int
}

func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	self.Offset = int(reader.ReadInt16())
}

/**
	运行时常量池的下标[由两字节操作数给出]
 */
type Index16Instruction struct {
	index	uint
}

func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.index = uint(reader.ReadUint8())
}
