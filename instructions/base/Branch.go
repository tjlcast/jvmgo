package base

import "jvmgo/rtda"

func Branch(frame *rtda.Frame, offset int) {
	pc := frame.NextPC()
	nextPC := pc + offset
	frame.SetNextPC(nextPC)
}
