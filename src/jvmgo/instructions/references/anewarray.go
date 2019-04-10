package references

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
)

//anewarray 指令用来创建引用类型数组
type ANEW_ARRAY struct{
	base.Index16Instruction


}
//anewarray 指令也需要两个操作数。第一个操作数是 uint16 索引，来自字节码。通过这个索引可以从
//当前类的运行时常量池中找到一个类符号引用，解析这个符号引用就可以得到数组元素的类。第二个操
//作数是数组长度，从操作数栈中弹出。



func (self *ANEW_ARRAY) Execute(frame *rtda.Frame) {

	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)

	componentClass := classRef.ResolvedClass()
	stack := frame.OperandStack()
	count := stack.PopInt()
	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}

	arrClass := componentClass.ArrayClass()
	arr := arrClass.NewArray(uint(count))
	stack.PushRef(arr)
}





