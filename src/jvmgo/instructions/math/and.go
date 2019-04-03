package math

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

//布尔运算指令只能操作 int 和 long 变量，分为按位与（ and ）、按位或（ or ）、按位异或（ xor ） 3 种。
//以按位与为例介绍布尔运算指令
type IAND struct {
	base.NoOperandsInstruction
}


func (i * IAND) Execute (frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 & v2
	stack.PushInt(result)
}

type LAND struct {
	base.NoOperandsInstruction
}


func (i * LAND) Execute (frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 & v2
	stack.PushLong(result)
}