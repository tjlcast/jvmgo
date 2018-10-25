package rtda

type Frame struct {
	lower			*Frame
	localVars		LocalVars
	operandStack	*OperandStack
}

/**
	new frame
 */
func newFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars:		newLocalVars(maxLocals),
		operandStack:	newOperandStack(maxStack),
	}
}
