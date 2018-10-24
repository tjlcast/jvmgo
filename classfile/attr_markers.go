package classfile

/**
	@doc:	标记类,支持空读
 */
type MarkerAttribute struct {

}

/**
	@doc:	表示该字段或方法不建议使用
 */
type DeprecatedAttribute struct {
	MarkerAttribute
}

/**
	@doc:	标记源文件中不存在，由编译器生产的类成员
 */
type SyntheticAttribute struct {
	MarkerAttribute
}

func (self *MarkerAttribute) readInfo(reader *ClassReader) {
	// read nothing.
}
