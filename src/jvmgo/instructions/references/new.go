package references

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
)

type NEW struct {
	base.Index16Instruction
}

func (self *NEW) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)

	class := classRef.ResolvedClass()

	if  class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationErr")
	}


	ref := class.NewObject()
	frame.OperandStack().PushRef(ref)
}