package stores

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)


//存储指令把变量从操作数栈顶弹出，然后存入局部变量表
type LSTORE struct {
	base.Index8Instruction
}

func (ins *LSTORE) Execute(frame *rtda.Frame) {
	_lstore(frame, ins.Index)
}

type LSTORE_0 struct {
	base.NoOperandsInstruction
}

func (ins *LSTORE_0) Execute(frame *rtda.Frame){
	_lstore(frame, 0)
}

type LSTORE_1 struct {
	base.NoOperandsInstruction
}

func (ins *LSTORE_1) Execute(frame *rtda.Frame){
	_lstore(frame, 1)
}

type LSTORE_2 struct {
	base.NoOperandsInstruction
}

func (ins *LSTORE_2) Execute(frame *rtda.Frame){
	_lstore(frame, 2)
}

type LSTORE_3 struct {
	base.NoOperandsInstruction
}

func (ins *LSTORE_3) Execute(frame *rtda.Frame){
	_lstore(frame, 3)
}

func _lstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index, val)
}