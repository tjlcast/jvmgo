package main

import (
	"jvmgo/classfile"
	"jvmgo/rtda"
	"fmt"
	"jvmgo/instructions/base"
	"jvmgo/instructions"
)

/**
	参数是 MemberInfo指针， 调用 MemberInfo 结构体的
	CodeAttribute() 方法可以获取它的Code属性
 */

func interpret(methodInfo *classfile.MemberInfo) {
	codeAttr := methodInfo.CodeAttribute()
	maxLocals := codeAttr.MaxLocals()
	maxStack := codeAttr.MaxStack()
	bytecode := codeAttr.Code()

	// 生成一个thread，并给 thread 添加一个 Frame，并压栈.
	thread := rtda.NewThread()
	frame := thread.NewFrame(maxLocals, maxStack)
	thread.PushFrame(frame)
	defer catchErr(frame)
	loop(thread, bytecode)
}

func catchErr(frame *rtda.Frame) {
	if r := recover(); r!=nil {
		fmt.Println("LocalVars: %v\n", frame.LocalVars())
		fmt.Println("OperandStack: %v\n", frame.OperandStack())
		panic(r)
	}
}

/**
	从 thread 的 stack 栈顶Frame拉出来，并在该 frame 上进行指令操作。
 */
func loop(thread *rtda.Thread, bytecode []byte) {

	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}

	for {
		pc := frame.NextPC()
		thread.SetPC(pc)

		// decode
		reader.Reset(bytecode, pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)

		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())

		// execute
		fmt.Println("pc %sd inst: %T %v\b", pc, inst, inst)
		inst.Execute(frame)
	}
}

