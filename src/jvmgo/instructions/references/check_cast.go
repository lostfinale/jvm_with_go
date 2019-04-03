package references

import (
	"jvmgo/rtda/heap"
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

//heckcast 指令和 instanceof 指令很像，区别在于： instanceof 指令会改变操作数栈（弹出对象引用，推
//入判断结果）； checkcast 则不改变操作数栈（如果判断失败，直接抛出 ClassCastException 异常）
type CHECK_CAST struct{
	base.Index16Instruction
}

func (self *CHECK_CAST) Execute(frame *rtda.Frame) {

	//先从操作数栈中弹出对象引用，再推回去，这样就不会改变操作数栈的状态
	stack := frame.OperandStack()
	ref := stack.PopRef()
	stack.PushRef(ref)
	if ref == nil {
		return
	}
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if !ref.IsInstanceOf(class) {
		panic("java.lang.ClassCastException")
	}
}