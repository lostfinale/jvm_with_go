package constants

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

// ldc 和 ldc_w 指令的区别仅在于操作数的宽度
type LDC struct{ base.Index8Instruction }
type LDC_W struct{ base.Index16Instruction }
type LDC2_W struct{ base.Index16Instruction }


func (self *LDC) Execute(frame *rtda.Frame) {
	_ldc(frame, self.Index)
}
func (self *LDC_W) Execute(frame *rtda.Frame) {
	_ldc(frame, self.Index)
}

func (self *LDC2_W) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	c := cp.GetConstant(self.Index)
	switch c.(type) {
	case int64: stack.PushLong(c.(int64))
	case float64: stack.PushDouble(c.(float64))
	default: panic("java.lang.ClassFormatError")
	}
}

func _ldc(frame *rtda.Frame, index uint) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	c := cp.GetConstant(index)
	switch c.(type) {
	case int32: stack.PushInt(c.(int32))
	case float32: stack.PushFloat(c.(float32))
		// case string:  在第8 章实现
		// case *heap.ClassRef:  在第9 章实现
	default: panic("todo: ldc!")
	}
}