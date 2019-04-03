package constants

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (ins * NOP) Execute(frame *rtda.Frame) {
	//Nothing to do.
}