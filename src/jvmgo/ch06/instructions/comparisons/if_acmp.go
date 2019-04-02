package comparisons

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

//if_acmpeq 和 if_acmpne 指令把栈顶的两个引用弹出，根据引用是否相同进行跳转


type IF_ACMPEQ struct {
	base.BranchInstruction
}


func (i *IF_ACMPEQ) Execute(frame *rtda.Frame) {
	if _acmp(frame) {
		base.Branch(frame, i.Offset)
	}
}

type IF_ACMPNE struct {
	base.BranchInstruction
}


func (self *IF_ACMPNE) Execute(frame *rtda.Frame) {
	if !_acmp(frame) {
		base.Branch(frame, self.Offset)
	}
}
func _acmp(frame *rtda.Frame) bool {
	stack := frame.OperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	return ref1 == ref2 // todo
}
