package stack

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

//swap 指令交换栈顶的两个变量
type SWAP struct {
	base.NoOperandsInstruction
}

func (s * SWAP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}