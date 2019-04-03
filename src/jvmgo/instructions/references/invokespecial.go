package references

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type INVOKE_SPECIAL struct{ base.Index16Instruction }
// hack!
func (self *INVOKE_SPECIAL) Execute(frame *rtda.Frame) {
	frame.OperandStack().PopRef()
}