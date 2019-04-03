package references

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda/heap"
	"jvmgo/rtda"
)

//instanceof 指令需要两个操作数
//第一个操作数是 uint16 索引，从方法的字节码中获取，通过这个索引
//可以从当前类的运行时常量池中找到一个类符号引用。第二个操作数是对象引用，从操作数栈中弹出。
type INSTANCE_OF struct {
	base.Index16Instruction


}

func (self *INSTANCE_OF) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	ref := stack.PopRef()
	if ref == nil {
		//如果是 null ，则把 0 推入操作数栈。用 Java 代码解释就是，如果引用 obj 是 null 的
		//话，不管 ClassYYY 是哪种类型，下面这条 if (obj instanceof ClassYYYY)... 判断都是 false
		stack.PushInt(0)
		return
	}
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if ref.IsInstanceOf(class) {
		stack.PushInt(1)
	} else {
		stack.PushInt(0)
	}
}