package extended

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

//根据引用是否是 null 进行跳转， ifnull 和 ifnonnull 指令把栈顶的引用弹出
type IFNULL struct {
	base.BranchInstruction
}

func (i *IFNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		base.Branch(frame, i.Offset)
	}
}

type IFNONNULL struct{ base.BranchInstruction }

func (self *IFNONNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		base.Branch(frame, self.Offset)
	}
}
