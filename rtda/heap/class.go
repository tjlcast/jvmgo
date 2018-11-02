package heap

import "jvmgo/classfile"

type Class struct {
	accessFlags			uint16
	name 				string	// thisClassName
	superClassName		string
	interfaceNames		[]string
	constantPool		*ConstantPool
	fields				[]*Field
	methods   			[]*Method
	loader				*ClassLoader
	superClass			*Class
	interfaces			[]*Class
	instanceSlotCount	uint
	staicSlotCount		uint
	staticVars			*Slot
}

func (self *Class) newClass(cf *classfile.ClassFile) *Class{
	class := &Class{}
	class.accessFlags = cf.AccessFlag()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())

	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())

	return class
}

func (self *Class) IsPublic() bool {
	return 0 != self.accessFlags & ACC_PUBLIC
}
