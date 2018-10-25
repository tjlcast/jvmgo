package base

type BytecodeReader struct {
	code		[]byte
	pc			int
}

func (self *BytecodeReader) Reset(code []byte, pc int) {
	self.code = code
	self.pc = pc
}

func (self *BytecodeReader) ReadUint8() uint8 {
	i := self.code[self.pc]
	self.pc++
	return i
}

func (self *BytecodeReader) ReadUint16() uint16 {
	byteHigh := uint16(self.ReadUint8())
	byteLow := uint16(self.ReadUint8())
	return (byteHigh<<8) | byteLow
}

func (self *BytecodeReader) ReadInt16() int16 {
	return int16(self.ReadUint16())
}

func (self *BytecodeReader) ReadInt32() int32 {
	byte1 := int32(self.ReadUint8())
	byte2 := int32(self.ReadUint8())
	byte3 := int32(self.ReadUint8())
	byte4 := int32(self.ReadUint8())

	return (byte1<<24) | (byte2<<16) | (byte3<<8) | (byte4)
}
