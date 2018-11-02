package constants

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

/**
	该系列指令吧隐含在操作码中的常量值推入 [操作数栈] 中
 */

// -------- null ---------
type ACONST_NULL struct {
	base.NoOperandInstruction
}

func (self *ACONST_NULL) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushRef(nil)
}


// -------- double 0.0 ---------
type DCONST_0 struct {
	base.NoOperandInstruction
}

func (self *DCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(0.0)
}


// -------- double 1.0 ---------
type DCONST_1 struct {
	base.NoOperandInstruction
}

func (self *DCONST_1) Execute(frame *rtda.Frame)  {
	frame.OperandStack().PushDouble(1.0)
}


// -------- Float 0.0 ---------
type FCONST_0 struct {
	base.NoOperandInstruction
}

func (self *FCONST_0) Execute(frame *rtda.Frame)  {
	frame.OperandStack().PushFloat(0.0)
}


// -------- Float 1.0 ---------
type FCONST_1 struct {
	base.NoOperandInstruction
}

func (self *FCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(1.0)
}


// -------- Float 2.0 ---------
type FCONST_2 struct {
	base.NoOperandInstruction
}

func (self *FCONST_2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(2.0)
}


// -------- Integer M1 ---------
type ICONST_M1 struct {
	base.NoOperandInstruction
}

func (self *ICONST_M1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(-1)
}


// -------- Integer 0 ---------
type ICONST_0 struct {
	base.NoOperandInstruction
}

func (self *ICONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(0)
}


// -------- Integer 1 ---------
type ICONST_1 struct {
	base.NoOperandInstruction
}

func (self *ICONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(1)
}


// -------- Integer 2 ---------
type ICONST_2 struct {
	base.NoOperandInstruction
}

func (self *ICONST_2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(2)
}


// -------- Integer 3 ---------
type ICONST_3 struct {
	base.NoOperandInstruction
}

func (self *ICONST_3) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(3)
}


// -------- Integer 4 ---------
type ICONST_4 struct {
	base.NoOperandInstruction
}

func (self *ICONST_4) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(4)
}


// -------- Integer 5 ---------
type ICONST_5 struct {
	base.NoOperandInstruction
}

func (self *ICONST_5) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(5)
}


// -------- Long 0  ---------
type LCONST_0 struct {
	base.NoOperandInstruction
}

func (self *LCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(0)
}


// -------- Long 1 ---------
type LCONST_1 struct {
	base.NoOperandInstruction
}

func (self *LCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(1)
}
