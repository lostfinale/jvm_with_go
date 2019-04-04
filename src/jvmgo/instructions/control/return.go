package control

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

//return指令
type RETURN struct{ base.NoOperandsInstruction } // Return void from method
func (self *RETURN) Execute(frame *rtda.Frame) {
	frame.Thread().PopFrame()
}


type ARETURN struct{ base.NoOperandsInstruction } // Return reference from method


func (self *ARETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	ref := currentFrame.OperandStack().PopRef()
	invokerFrame.OperandStack().PushRef(ref)
}

type DRETURN struct{ base.NoOperandsInstruction } // Return double from method


func (self *DRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopDouble()
	invokerFrame.OperandStack().PushDouble(val)
}

type FRETURN struct{ base.NoOperandsInstruction } // Return float from method


func (self *FRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopFloat()
	invokerFrame.OperandStack().PushFloat(val)
}

type IRETURN struct{ base.NoOperandsInstruction } // Return int from method
func (self *IRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	//当前帧
	currentFrame := thread.PopFrame()
	//上一帧
	invokerFrame := thread.TopFrame()
	//弹出当前帧的处理结果
	retVal := currentFrame.OperandStack().PopInt()
	//放入上一帧
	invokerFrame.OperandStack().PushInt(retVal)
}

type LRETURN struct{ base.NoOperandsInstruction } // Return long from method

func (self *LRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopLong()
	invokerFrame.OperandStack().PushLong(val)
}