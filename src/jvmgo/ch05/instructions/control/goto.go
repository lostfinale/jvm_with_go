package control

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type GOTO struct {
	base.BranchInstruction
}

//goto 指令进行无条件跳转
func (g *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, g.Offset)
}
