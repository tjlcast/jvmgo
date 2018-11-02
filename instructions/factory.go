package instructions

import "jvmgo/instructions/base"

var (
	nop				= &NOP{}
	aconst_null		= &ACOUST_NULL{}
)

func NewInstruction(opcode byte) base.Instruction {
	switch opcode {
	case 0x00:
		return &NOP{}
	}
}
