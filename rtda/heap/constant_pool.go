package heap

import "jvmgo/classfile"

type Constant interface {}

type ConstantPool struct {
	class 			*Class
	consts 			[]Constant
}

func newConstantPool(class *Class, cfCp classfile.ConstantPool) *ConstantPool {
	// todo
	return nil
}

func (self *ConstantPool) GetConstant(index uint) Constant {
	// todo
	return nil
}
