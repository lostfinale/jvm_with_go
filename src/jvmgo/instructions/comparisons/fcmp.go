package comparisons

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

//fcmpg 和 fcmpl 指令用于比较 float 变量

type FCMPG struct {
	base.NoOperandsInstruction
}


func (f *FCMPG) Execute(frame *rtda.Frame) {
	_fcmp(frame, true)
}

type FCMPL struct {
	base.NoOperandsInstruction
}

func (f *FCMPL) Execute(frame *rtda.Frame) {
	_fcmp(frame, false)
}

//当两个 float 变量中至少有一个是 NaN 时，用 fcmpg 指令比较的结果是 1 ，
//而用 fcmpl 指令比较的结果是 -1
func _fcmp(frame *rtda.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 < v1 {
		stack.PushInt(-1)
	} else if gFlag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}

}