package classfile

/**
	Code_attribute {
		u2 attribute_name_index;
		u4 attreibute_length;
		u2 maxx_locals;
		u4 code_length;
		u1 code[code_length]
		u2 exception_table_length;
		{

		} exception_table[exception_table_length];
		u2 attributes_count;
		attribute_info attributes[attributes_count];
	}
 */

/**
	Code 是变长属性，存在于 method_info结构中
	Code 属性中存放字节码等方法相关信息
 */

type CodeAttribute struct {
	cp			ConstantPool						// 本身是数组
	maxStack	uint16								// 操作数栈的最大深度
	maxLocals	uint16								// 给出局部变量表大小
	code		[]byte								//
	exceptionTable		[]*ExceptionTableEntry
	attributes  		[]AttributeInfo
}

/**
	这里是异常结构体，用于描述结构。一般存在Code Attr 中
 */
type ExceptionTableEntry struct {
	startPc			uint16
	endPc			uint16
	handlerPc		uint16
	catchType		uint16
}

//======= function ========

func (self *CodeAttribute) ConstantPool() ConstantPool{
	return self.cp
}

func (self *CodeAttribute) MaxStack() uint {
	return uint(self.maxStack)
}

func (self *CodeAttribute) MaxLocals() uint {
	return uint(self.maxLocals)
}

func (self *CodeAttribute) Code() []byte {
	return self.code
}

func (self *CodeAttribute) readInfo(reader *ClassReader) {
	self.maxStack = reader.readUint16()
	self.maxLocals = reader.readUint16()
	codeLength := reader.readUint32()
	self.code = reader.readBytes(codeLength)
	self.exceptionTable = readExceptionTable(reader)
	self.attributes = readAttributes(reader, self.cp)
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry{
	exceptionTableLength := reader.readUint16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)

	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc:	reader.readUint16(),
			endPc:		reader.readUint16(),
			handlerPc:	reader.readUint16(),
			catchType:	reader.readUint16(),
		}
	}

	return exceptionTable
}
