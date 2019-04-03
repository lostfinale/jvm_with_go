package references

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
)

type GET_STATIC struct {

	base.Index16Instruction

}


func (self *GET_STATIC) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)
	field := fieldRef.ResolveField()
	class := field.Class()
	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := class.StaticVars()
	stack := frame.OperandStack()
	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I': stack.PushInt(slots.GetInt(slotId))
	case 'F': stack.PushFloat(slots.GetFloat(slotId))
	case 'J': stack.PushLong(slots.GetLong(slotId))
	case 'D': stack.PushDouble(slots.GetDouble(slotId))
	case 'L', '[': stack.PushRef(slots.GetRef(slotId))
	}
}