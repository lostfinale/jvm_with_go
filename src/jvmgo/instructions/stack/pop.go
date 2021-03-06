package stack

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

//用于弹出 int 、 float 等占用一个操作数栈位置的变量
type POP struct {
	base.NoOperandsInstruction
}

func (p *POP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

// double 和 long 变量在操作数栈中占据两个位置，需要使用 pop2 指令弹出
type POP2 struct {
	base.NoOperandsInstruction
}

func (p *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
