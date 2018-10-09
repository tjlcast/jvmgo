package classfile

type ClassFile struct {
	// magic			uint32
	minorVersion		uint16
	majorVersion		uint16
	constantPool		ConstantPool
	accessFlags			uint16
	thisClass			uint16
	superClass			uint16
	interfaces			[]uint16
	fields				[]*MemberInfo
	methods				[]*MemberInfo
	attribures 			[]AttributeInfo
}

func Parse(classData []byte) (cf *ClassFile, err error) {
	// todo
}

func (self *ClassFile) read(reader *ClassReader) {
	// todo
}

func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	// todo
}

func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	// todo
}

func (self *ClassFile) MinorVersion() uint16 {
	// todo
}

func (self *ClassFile) MajorVersion() uint16 {
	// todo
}

func (self *ClassFile) ConstantPool() ConstantPool {
	// todo
}

func (self *ClassFile) AccessFlag() uint16 {
	// todo
}

func (self *ClassFile) Fields() []*MemberInfo {
	// todo
}

func (self *ClassFile) Method() []*MemberInfo {
	// todo
}

func (self *ClassFile) ClassName() string {
	// todo
}

func (self *ClassFile) SuperClassName() string {
	// todo
}

func (self *ClassFile) InterfaceNames() []string {
	// todo
}

