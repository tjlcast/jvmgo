package math

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
	"math"
)

/**
	算法指令：
		加法	add
		相减	sub
		乘法	mul
		除法	div
		求余	rem
		相反	reg
 */


/**
   先从操作栈中弹出两个 long 变量，再进行求余操作，然后把结果推入操作数栈。
*/
type DREM struct {
	base.NoOperandInstruction
}

func (self *DREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	result := math.Mod(v1, v2)
	stack.PushDouble(result)
}

type FREM struct {
	base.NoOperandInstruction
}

func (self *FREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	//result := v1 % v2
	result := math.Float32bits()
	stack.PushFloat(result)
}

/**
	先从操作栈中弹出两个 int 变量，再进行求余操作，然后把结果推入操作数栈。
 */

type IREM struct {
	base.NoOperandInstruction
}

func (self *IREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	result := v1 % v2
	stack.PushInt(result)
}

type LREM struct {
	base.NoOperandInstruction
}

