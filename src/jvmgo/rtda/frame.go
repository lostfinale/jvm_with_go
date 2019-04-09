package rtda

import "jvmgo/rtda/heap"

//帧
type Frame struct {
	lower *Frame
	localVars LocalVars
	operandStack *OperandStack
	thread       *Thread
	method       *heap.Method
	nextPC int
}
/*
func NewFrame(maxLocals, maxStack uint) *Frame{
	return &Frame {
		localVars:newLocalVars(maxLocals),
		operandStack:newOperandStack(maxStack),
	}
}
*/

func newFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame{
		thread: thread,
		method: method,
		localVars: newLocalVars(method.MaxLocals()),
		operandStack: newOperandStack(method.MaxStack()),
	}
}

//先判断类的初始化是否已经开始，如果还
//没有，则需要调用类的初始化方法，并终止指令执行。但是由于此时指令已经执行到了一半，也就是说
//当前帧的 nextPC 字段已经指向下一条指令，所以需要修改 nextPC ，让它重新指向当前指令
func (self *Frame) RevertNextPC() {
	self.nextPC = self.thread.pc
}


func (f *Frame) LocalVars() LocalVars {
	return f.localVars
}
func (f *Frame) OperandStack() *OperandStack {
	return f.operandStack
}

func (f *Frame) Thread() *Thread {
	return f.thread
}
func (f *Frame) NextPC() int {
	return f.nextPC
}
func (f *Frame) SetNextPC(nextPC int) {
	f.nextPC = nextPC
}

func (f *Frame) Method() *heap.Method {
	return f.method
}

