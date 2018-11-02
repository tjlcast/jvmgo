package math

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

/**
	int 左移动
 */

type ISHL struct {
	base.NoOperandInstruction
}

func (self *ISHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f // 对于32位的数，有且只能移动32
	result := v1 << s
	stack.PushInt(result)
}

/**
	int 右移动
 */

type ISHR struct {
	base.NoOperandInstruction
}

/**
	int 逻辑右移动
 */

type IUSHR struct {
	base.NoOperandInstruction
}

func (self *IUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := int32(uint32(v1) >> s)
	stack.PushInt(result)
}

/**
	long 左移动
 */

type LSHL struct {
	base.NoOperandInstruction
}

/**
	long 右移动
 */

type LSHR struct {
	base.NoOperandInstruction
}

/**
	long 逻辑右移动
 */

type LUSHR struct {
	base.NoOperandInstruction
}


