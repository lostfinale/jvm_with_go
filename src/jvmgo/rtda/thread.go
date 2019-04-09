package rtda

import "jvmgo/rtda/heap"

type Thread struct {
	pc int //pc寄存器
	stack *Stack //java虚拟机栈
}


/*
JVM
  Thread
    pc
    Stack
      Frame
        LocalVars
        OperandStack
*/

func NewThread() *Thread {
	return &Thread{stack:newStack(1024)}
}

//入栈
func (t *Thread) PushFrame(frame *Frame) {
	t.stack.push(frame)
}

//出栈
func (t *Thread) PopFrame() *Frame {
	return t.stack.pop()
}

//当前帧
func (t *Thread) CurrentFrame() *Frame {
	return t.stack.top()
}


//类似当前帧
func (t *Thread) TopFrame() *Frame {
	return t.stack.top()
}


func (t *Thread) PC() int {
	return t.pc
}
func (t *Thread) SetPC(pc int) {
	t.pc = pc
}

func (t *Thread) NewFrame(method *heap.Method) *Frame {
	return newFrame(t, method)
}

func (t *Thread) IsStackEmpty() bool {
	return t.stack.isEmpty()
}