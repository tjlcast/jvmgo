package classfile

/**
	定长，出现在 field_info 结构中，用于表示常量表达式的值
 */

type ConstantValueAttribute struct {
	constantValueIndex		uint16
}

func (self *ConstantValueAttribute) readInfo(reader *ClassReader) {
	self.constantValueIndex = reader.readUint16()
}

func (self *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return self.constantValueIndex
}
