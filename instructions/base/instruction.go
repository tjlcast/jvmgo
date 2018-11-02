package base

import "jvmgo/rtda"

/**
	指令的一级对象
 */
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
	指令的二级对象，表示没有操作数的指令
 */
type NoOperandInstruction struct {}

func (self *NoOperandInstruction) FetchOperands(reader *BytecodeReader) {
	// nothing to do
}


/**
	指令的二级指令，表示一个跳转指令
	Offset 存放跳转偏移量
 */
type BranchInstruction struct {
	Offset			int
}

func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	self.Offset = int(reader.ReadInt16())
}


/**
	指令的二级指令，存储和加载指令需要根据索引存取局部变量表
	索引由单字节给出
 */
type Index8Instruction struct {
	Index 			uint
}

func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
}


/**
	指令的二级指令，有的指令需要访问运行时常量池，该常量池的索引由2字节给出
 */
type Index16Instruction struct {
	index	uint
}

func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.index = uint(reader.ReadUint16())
}
