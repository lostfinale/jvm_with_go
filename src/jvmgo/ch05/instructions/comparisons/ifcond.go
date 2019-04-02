package comparisons

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

//
/*
	if<cond> 指令把操作数栈顶的 int 变量弹出，然后跟 0 进行比较，满足条件则跳转。
	假设从栈顶弹出的变量是 x ，则指令执行跳转操作的条件如下：
	·ifeq ： x==0
	·ifne ： x ！ =0
	·iflt ： x<0
	·ifle ： x<=0
	·ifgt ： x>0
	·ifge ： x>=0



 */
type IFEQ struct {
	base.BranchInstruction
}

func (self *IFEQ) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val == 0 {
		base.Branch(frame, self.Offset)
	}
}

type IFNE struct {
	base.BranchInstruction
}


func (self *IFNE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val != 0 {
		base.Branch(frame, self.Offset)
	}
}


type IFLT struct {
	base.BranchInstruction
}

func (self *IFLT) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val < 0 {
		base.Branch(frame, self.Offset)
	}
}

type IFLE struct {
	base.BranchInstruction
}

func (self *IFLE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val <= 0 {
		base.Branch(frame, self.Offset)
	}
}

type IFGT struct {
	base.BranchInstruction
}

func (self *IFGT) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val > 0 {
		base.Branch(frame, self.Offset)
	}
}

type IFGE struct {
	base.BranchInstruction
}


func (self *IFGE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val >= 0 {
		base.Branch(frame, self.Offset)
	}
}